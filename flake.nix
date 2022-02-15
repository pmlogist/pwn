{
  description = "Pown my app";

  inputs.flake-utils.url = "github:numtide/flake-utils";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ ];
        };

        go = pkgs.go_1_16;

      in rec {
        devShell = pkgs.mkShell {
          buildInputs = [ go ];
          shellHook = ''
            export TMPDIR=/tmp
            export GOPATH=$TMPDIR
            export PATH=$PATH:$(go env GOPATH)/bin 

            curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
          '';
        };
      });
}
