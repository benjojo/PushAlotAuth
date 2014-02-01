PushAlotAuth
============

A service that will send you a "Pushalot" notification when any kind of auth happens on a Linux system.

![Say it with pixels](http://i.imgur.com/FVyjd5b.png)

## Don't want to go though the setup with go?

I can relate with that. Just download a precompiled version for your OS/Arch in the releases and run like so

`./PushAlotAuth`

After it makes you a first time config, edit that config and test that it works by running it again.

After that you can run it "forever" by doing

```bash
nohup ./PushAlotAuth &
```

##Sample setup

```json
{
  "DeviceToken": "DevToken",
  "Application": "Apptoken",
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
