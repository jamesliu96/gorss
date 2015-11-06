package gorss

import (
    "testing"
)

func Test(*testing.T) {
    c := &Channel{
        Title: "test channel",
        Link: "http://example.com/testchannel",
        Description: "this is test channel",
    }

    c.AddItem(&Item{
        Title: "hello world",
        Link: "http://example.com/path/to/helloworld",
        Description: "foo bar",
    })

    c.Publish()
    c.PublishCompressed()
}
