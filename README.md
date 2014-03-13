PushOverAuth
============

A service that will send you a "PushOver" notification when any kind of auth happens on a Linux system.

![Say it with pixels](http://i.imgur.com/5VSbFC9.png)

## Don't want to go though the setup with go?

I can relate with that. Just download a precompiled version for your OS/Arch in the releases and run like so

`./PushOverAuth`

After it makes you a first time config, edit that config and test that it works by running it again.

After that you can run it "forever" by doing

```bash
nohup ./PushOverAuth &
```

##Sample setup

```json
{
  "UserToken": "Usertoken",
  "AppToken": "Apptoken",
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
