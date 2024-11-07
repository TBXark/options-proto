# options-proto

This is a proto file that defines extra options for rpc methods.

### Example
```protobuf
syntax = "proto3";

package bot.v1;

import "tbxark/options/annotations.proto";

service CounterService {
  rpc Start(StartRequest) returns (StartResponse) {
    option (tbxark.options.options) = {
      key: "bot",
      extra: [
        {
          key: "command",
          value: "start",
        }
      ]
    };
  }
  rpc Counter(CounterRequest) returns (CounterResponse) {
    option (tbxark.options.options) = {
      key: "bot",
      extra: [
        {
          key: "command",
          value: "count",
        },
        {
          key: "callback_query",
          value: "count",
        }
      ]
    };
  }
  rpc Unknown(UnknownRequest) returns (UnknownResponse);
}

...
```
