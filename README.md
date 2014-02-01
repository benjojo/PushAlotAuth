PushAlotAuth
============

A service that will send you a "Pushalot" notification when any kind of auth happens on a Linux system.

##Sample setup

```json
{
  "Token": "Fillmein",
  "Watches": [
    {
      "Path": "/var/log/auth.log",
      "TriggerWords": [
        "Accepted publickey",
        "Accepted password"
      ]
    }
  ]
}
```