What is Agha Balasar?
================

It's a health-check service(written in Go) which sends request to your http services every *n* seconds to insure everything is fine.

How to Use?
================

Create a YAML file that contains your services properties:
```yaml
services:
  - name: google
    request-url: https://www.google.com
    period: 10
    requests-count: 5
    failures-count: 3
    timeout: 1.5
    post: false
    slacks:
      - https://hooks.slack.com/services/blablabla

  - name: bing
    request-url: https://www.bing.com
    period: 20
    requests-count: 4
    failures-count: 5
    timeout: 2
    post: false
    webhooks:
      - https://www.mywebsite.com/my/webhook/url
```

License
================

Licensed under the GNU General Public License v3.0.

Author
================

Arya Hadi (arya.hadi97@gmail.com)