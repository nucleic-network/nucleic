# x/rollapp

## Abstract

RollApps are standardized blockchains that are easy to bootstrap and provide a seamless user experience. RollApps come with pre-defined modules that expediate the development process, incorporating features such as minting native tokens, onchain governance, out-of-the-box bridging and much more.

## Contents

1. **[Concepts](#concepts)**
2. **[State](#state)**
3. **[Events](#events)**
4. **[Keepers](#keepers)**
5. **[Hooks](#hooks)**
6. **[Queries](#queries)**

## Concepts

This document specifies the rollapp module.

This module allows rollapp register onto Nucleic chain.

If a rollapp is registered on the Nucleic chain, sequencers should be able to nominate themselves for the rollapp by running a rollapp node.

The Nucleic chain selects a sequencer for a block n blocks before. This is done through leader rotation based on restaked TIA signaled to each sequencer on that rollapp.

## State

Each RollApp is identified by a `rollapp_id` folows the same standard as cosmos `chain_id`.

Here is a struct of a RollApp:

```protobuf
message Rollapp {
  string rollapp_id = 1;
  string creator = 2;
  // version is the software and configuration version.
  // starts from 1 and increases by one on every MsgUpdateState
  uint64 version = 3;
  // max_sequencers is the maximum number of sequencers.
  uint64 max_sequencers = 4;
  // permissioned_addresses is a bech32-encoded address list of the sequencers
  // that are allowed to serve this rollapp_id. In the case of an empty list, the
  // rollapp is considered permissionless.
  repeated string permissioned_addresses = 5;
  // token_metadata is a list of TokenMetadata that are registered on this
  // rollapp
  repeated TokenMetadata token_metadata = 6;
  // genesis_state is a partial repr of the state the hub can expect the rollapp
  // to be in upon genesis
  RollappGenesisState genesis_state = 7 [ (gogoproto.nullable) = false ];
  // channel_id will be set to the canonical IBC channel of the rollapp.
  string channel_id = 8;
  // frozen is a boolean that indicates if the rollapp is frozen.
  bool frozen = 9;
}
```

## Events

## Keepers

## Hooks

## Queries
