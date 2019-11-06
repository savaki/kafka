[![GoDoc](https://godoc.org/github.com/savaki/kafka?status.svg)](https://godoc.org/github.com/savaki/kafka)

# kafka

`kafka` is a pure go Kafka library.

***Warning*** - this library is under heavy construction right now.  Definitely NOT safe for use.

## Motivations

The `kafka` library is intended as a base library to write Kafka tooling.  As its intended
to be used for writing more than just Kafka consumers and producers, `kafka` needs to 
support the entire Kafka protocol. 
to support the entire range of the [Kafka protocol](https://kafka.apache.org/protocol).   

Go has a number of high quality Kafka libraries.  Unfortunately, as of today, none of them 
met our requirements.  The libraries evaluated were:

* [sarama](https://shopify.github.io/sarama/) - An excellent Kafka library written by folks
at Shopify.  While sarama supports a significant portion of the Kafka protocol, it doesn't
support the entire protocol.  Additionally, as the `kafka-go` has noted, it passes all values
as pointers which causes a large number of dynamic memory allocations and more frequent
garbage collections.   

* [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) - Another excellent
Kafka library by Confluent itself.  Like `sarama`, `confluent-kafka-go` does not support the
full Kafka protocol.  In addition, `confluent-kafka-go` is a wrapper based around 
[librdkafka](https://github.com/edenhill/librdkafka) which means it uses cgo.  And to quote
Dave Cheney, [cgo is not Go](https://dave.cheney.net/2016/01/18/cgo-is-not-go)

* [kafka-go](https://github.com/segmentio/kafka-go) - To round out the set of libraries we
looked at is `kafka-go` by Segment.  Another great package with a simple to use API.  Lots of
things to like here.  The only real issue is that the entire Kafka protocol is not supported.
And from talking with the team, I know that supporting the entire Kafka protocol isn't their
intent.

## Kafka Versions

Most of the `kafka` library is generated from the json protocol definition.  In addition, 
the library dynamically negotiates the version for each api key so it should (in theory)
support any kafka broker that allows for protocol version negotiation.

## Go version

`kafka` - For now, `kafka` requires golang 1.13 or better as it makes use of the updated
errors handling.  
