# Verifiable Data Registry Core

## Introduction

Verifiable Data Registry provided an interface for storing Decentralised Identifiers (DID documents).
The library uses plugin-based architecture, using it as a dependency requires additional steps.

## Building

Here is the [README.md](https://gitlab.eclipse.org/eclipse/xfsc/dev-ops/building/go-plugin/-/blob/main/README.md#building-go-services-with-plugin-based-dependencies) describing the specifics of build process for services, where the dependency is used.

## Usage

The core module defines an interface VerifiableDataRegistry that defines methods implemented in each plugin. 
The core module loads a plugin from specified location and exports the VerifiableDataRegistry implementation.

## Configuration

Following environment variables are required:
1. PLUGINS_LOCATION - path to folder that includes compiled plugin files
2. VDR_TYPE - name of compiled plugin file

