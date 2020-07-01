## TODO:

### Fixes and architectural improvements
- [x] Drop any Conf reference from clightd wrapper
- [x] Rework Clightd/Idle wrapper to only wrap calls
- [x] Move Clightd wrapper to its own external module
- [x] Fix: only start gamma once a first location has been received
- [x] Fix: gamma sunrise/sunset times are off by a couple of hours (UTC delta)
- [x] Support multiple ac states
- [ ] Move to single package only (drop all subdirs and move everything to main package)

### New features
- [ ] Implement conf parsing (https://github.com/spf13/viper)
- [ ] Implement dbus server exposing state
