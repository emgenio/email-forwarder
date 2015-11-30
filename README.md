# Email-forwarder
A microservice consumer written in go to forward incoming messages from a Rabbit MQ message queue.

## How it works
Email-forwarder instanciates a Rabbit MQ queue specified in a config yaml file, consume all messages that have been push into the queue and finally forward them to their destination via [Mandrill](https://www.mandrill.com/) (an email delivery api from Mailchimp).

## How to install

1. Clone the git repository:
  ```
  git clone https://github.com/emgenio/email-poller && cd email-poller
  ```
2. Set up your configurations in config.yaml
3. Compile and execute:
  ```
  make && ./build/worker
  ```

## Licence
MIT