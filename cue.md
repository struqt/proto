```sh
  cue import proto -I "$(brew --prefix)/opt/protobuf@29/include" -f demo/struqt/demo/v3/messages.proto
```

```sh
  cue vet -c demo/struqt/demo/v3/*.cue
```

---

```sh
  cue import proto -I "$(brew --prefix)/opt/protobuf@29/include" \
    -I googleapis -I grpc-gateway -I common -I demo -f demo/**/*.proto
```
