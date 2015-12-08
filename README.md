# Email-forwarder
A microservice consumer written in go to forward incoming messages from a Rabbit MQ message queue.

## How it works
Email-forwarder instanciates a Rabbit MQ queue specified in a config yaml file, consume all messages that have been push into the queue and finally forward them to their destination via [Mandrill](https://www.mandrill.com/) (an email delivery api from Mailchimp).

## How to install

1. Get the package: `$> go get github.com/emgenio/email-forwarder`
2. Run: `$> email-forwarder`

## Licence
MIT