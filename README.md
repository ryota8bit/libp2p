# libp2p
## set up
```bash
make deps
```

## build
```bash
make bld
```

## run
```bash
$ p2p-app -l 8080
2019/03/21 17:24:29 I am /ip4/127.0.0.1/tcp/8080/ipfs/QmeBCNL1Ap1fxC56FJdmQSfePibDXaSyGvVnE3CynqMUEz
2019/03/21 17:24:29 Now run "p2p-app -l 8081 -d /ip4/127.0.0.1/tcp/8080/ipfs/QmeBCNL1Ap1fxC56FJdmQSfePibDXaSyGvVnE3CynqMUEz" on a different terminal
2019/03/21 17:24:29 listening for connections
```

```bash
$ p2p-app -l 8081 -d /ip4/127.0.0.1/tcp/8080/ipfs/QmeBCNL1Ap1fxC56FJdmQSfePibDXaSyGvVnE3CynqMUEz
2019/03/21 17:25:16 I am /ip4/127.0.0.1/tcp/8081/ipfs/QmQH1d1JEZTPZEfo1mKRG4w6VFg1mHihhUMzFBYCPsP1Nn
2019/03/21 17:25:16 Now run "p2p-app -l 8082 -d /ip4/127.0.0.1/tcp/8081/ipfs/QmQH1d1JEZTPZEfo1mKRG4w6VFg1mHihhUMzFBYCPsP1Nn" on a different terminal
2019/03/21 17:25:16 opening stream
2019/03/21 17:25:16 We are  [QmQH1d1JEZTPZEfo1mKRG4w6VFg1mHihhUMzFBYCPsP1Nn QmeBCNL1Ap1fxC56FJdmQSfePibDXaSyGvVnE3CynqMUEz]
```