{
  description = "fdate";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = {
    self,
      nixpkgs,
  }: let
    projectName = "fdate";
    supportedSystems = ["x86_64-linux" "aarch64-darwin" "aarch64-linux" "x86_64-darwin" ];
    forSystems = systems: f:
      nixpkgs.lib.genAttrs systems
        (system: f system (import nixpkgs {inherit system;}));
    forAllSystems = forSystems supportedSystems;
  in {
    packages = forAllSystems (system: pkgs: {
      ${projectName} = pkgs.buildGoModule {
        pname = projectName;
        version = "1.1";
        vendorHash = "sha256-sUvYnu2GFhSgcK98NoFuN7v6bz3biA8d7Ug98YhcDIo=";
        src = ./.;
      };
      default = self.packages.${system}.${projectName};
    });

    formatter = forAllSystems (system: pkgs: pkgs.alejandra);

    devShells = forAllSystems (system: pkgs: {
      default = pkgs.mkShell {
        buildInputs = with pkgs; [ go ];
      };
    });
  };
}
