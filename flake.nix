{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = nixpkgs.legacyPackages.${system};
      in {
        devShells.default = pkgs.mkShell {
          packages = [
            # Languages
            pkgs.dotnet-sdk
            pkgs.go
            pkgs.nodejs
            pkgs.python3
            pkgs.typescript
            pkgs.yarn
            # Pulumi packages
            (pkgs.pulumi.withPackages (ps:
              with ps; [
                pulumi-go
                pulumi-python
                pulumi-nodejs
              ]))
            pkgs.pulumictl
            pkgs.delve # Go debugger
            pkgs.grpc
          ];
          shellHook = ''
            export LD_LIBRARY_PATH=${pkgs.lib.makeLibraryPath [pkgs.stdenv.cc.cc]}
            export GOPATH=$(pwd)/.go
            export PULUMI_HOME=$(pwd)/.pulumi
          '';
        };
      }
    );
}
