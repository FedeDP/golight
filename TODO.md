## TODO:

### Fixes and architectural improvements
- [x] Drop any Conf reference from clightd wrapper
- [x] Rework Clightd/Idle wrapper to only wrap calls
- [x] Move Clightd wrapper to its own external module
- [x] Fix: only start gamma once a first location has been received
- [x] Fix: gamma sunrise/sunset times are off by a couple of hours (UTC delta)
- [x] Support multiple ac states
- [x] Use a goeish project layout
- [x] Use go fmt on source tree

### New features
- [ ] Implement conf parsing (https://github.com/spf13/viper)
- [ ] Implement dbus server exposing state
- [ ] Implement CLI

### Improvements
- [ ] Add documentation to methods and functions
- [ ] Add some tests (at least for pkg/go-clightd)
- [ ] Add CI with tests

