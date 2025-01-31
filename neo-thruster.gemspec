require_relative "lib/thruster/version"

Gem::Specification.new do |s|
  s.name        = "neo-thruster"
  s.version     = Thruster::VERSION
  s.summary     = "Zero-config HTTP/2 proxy"
  s.description = "A fork of thruster - A zero-config HTTP/2 proxy for lightweight production deployments with zstd support"
  s.authors     = [ "Kevin McConnell", "Shivam Mishra" ]
  s.email       = "hey@shivam.dev"
  s.homepage    = "https://github.com/scmmishra/neo-thruster"
  s.license     = "MIT"

  s.metadata = {
    "homepage_uri" => s.homepage,
    "rubygems_mfa_required" => "true"
  }

  s.files = Dir[ "{lib}/**/*", "MIT-LICENSE", "README.md" ]
  s.bindir = "exe"
  s.executables << "thrust"
end
