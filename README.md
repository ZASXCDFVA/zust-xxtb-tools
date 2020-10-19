# zust-xxtb-tools

### Usage
```bash
./zust-xxtb-tools /path/to/configuration
```

### Configuration
```yaml
# auto push checkin log to telegram
telegram:
  token: "<bot token>"
  chat-id: "<chat id>"

default-location: &default-location
  province: "浙江省"
  city: "***"
  country: "***"

users:
- id: "0000000000"
  <<: *default-location
- id: "0000000001"
  province: "浙江省"
  city: "***"
  country: "***"
```