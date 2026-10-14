{
  description = "hacktober CTF app";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-26.05";
    home-manager.url = "github:nix-community/home-manager/release-26.05";
    home-manager.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = { nixpkgs, home-manager, ... }:
  let
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};

    hacktober = pkgs.buildGoModule.override { go = pkgs.go_1_27; } {
        pname = "hacktober";
        version = "0.1.0";
        src = ./.;
        vendorHash = "sha256-LhBbKXYmqY2vurRbYg7kcIlPQXf5V54Q9gccPGubEls=";
        nativeBuildInputs = [ pkgs.pkg-config ];
        buildInputs = [ pkgs.gpgme pkgs.btrfs-progs pkgs.lvm2 ];
    };

    challengeData = pkgs.runCommand "hacktober-challenges" { } ''
        cp -r ${./challenges} $out
    '';

    config = (pkgs.formats.toml { }).generate "hacktober.toml" {
        ssh = {
            Host = "0.0.0.0";
            Port = "22222";
            PublicUrl = "ssh://hacktober";
            KeyPath = "/persist/app/ssh_host_ed25519";
        };
        container = {
            ContainerPort = 1337;
            TimeoutInMinutes = 30;
            RegistryPrefix = "10.0.0.1:5000";
        };
        challenge.BaseDir = "${challengeData}/2024";
    };

    hostKey = pkgs.writeShellScript "hacktober-hostkey" ''
        key=/persist/app/ssh_host_ed25519
        [ -f "$key" ] || ${pkgs.openssh}/bin/ssh-keygen -t ed25519 -f "$key" -N "" -C hacktober
    '';
      in
    {
        packages.${system}.default = hacktober;

        homeConfigurations.app = home-manager.lib.homeManagerConfiguration {
            inherit pkgs;
            modules = [{
                home.username = "app";
                home.homeDirectory = "/home/app";
                home.stateVersion = "26.05";

                home.file.".config/hacktober.toml".source = config;

                systemd.user.startServices = "sd-switch";

                systemd.user.services.hacktober-hostkey = {
                    Unit.Description = "Generate hacktober SSH host key";
                    Service = { Type = "oneshot"; ExecStart = "${hostKey}"; };
                    Install.WantedBy = [ "default.target" ];
                };

                systemd.user.services.hacktober = {
                    Unit = {
                        Description = "hacktober SSH app";
                        Wants = [ "podman.socket" ];
                        After = [ "podman.socket" "hacktober-hostkey.service" ];
                        Requires = [ "hacktober-hostkey.service" ];
                    };
                    Service = {
                        ExecStart = "${hacktober}/bin/hacktober serve";
                        Environment = [ "PATH=/run/wrappers/bin:/run/current-system/sw/bin" ];
                        Restart = "on-failure";
                    };
                    Install.WantedBy = [ "default.target" ];
                };
            }];
        };
    };
}
