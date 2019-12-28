# Slack Mailer
Linux sendmail MTA compatible binary that forwards all emails as notifications to your Slack channel.
Its main purpose is for letting you know of failing cron jobs, unattended updates going wrong etc.

## Building and running
1. Install the app on slack to get the webhook url
2. Run `go build` to produce `slack-mailer` binary
3. Place the binary in the location of current sendmail binary (e.g. `/usr/sbin/sendmail`) and make it executable
4. Save following config to file named `slack-mailer.json` in either user home directory, /etc or /root:
```json
{
  "debug": false,
  "webhook_url": "YOUR_SLACK_WEBHOOK_URL"
}
```

> You can test by running: `echo "Hello Root" | mail -s "Test email" root`