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
```

License
================

Licensed under the GNU General Public License v3.0.

Author
================

Arya Hadi (arya.hadi97@gmail.com)