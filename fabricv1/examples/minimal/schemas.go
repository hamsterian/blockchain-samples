package main


	import (
		"github.com/hyperledger/fabric/core/chaincode/shim"
		iot "github.com/ibm-watson-iot/blockchain-samples/fabricv1/platform"
)

var schemas = `

{
    "API": {
        "createAsset": {
            "description": "Creates a new asset by class",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "asset": {
                                "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                "properties": {
                                    "assetID": {
                                        "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                        "type": "string"
                                    },
                                    "carrier": {
                                        "description": "The carrier in possession of this asset",
                                        "type": "string"
                                    },
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
                                            },
                                            "location": {
                                                "description": "A geographical coordinate",
                                                "properties": {
                                                    "latitude": {
                                                        "type": "number"
                                                    },
                                                    "longitude": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "temperature": {
                                        "description": "Temperature of an asset's contents in degrees Celsuis",
                                        "type": "number"
                                    }
                                },
                                "required": [
                                    "assetID"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "createAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssets": {
            "description": "Delete all assets from world state, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAllAssets"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAsset": {
            "description": "Delete an asset from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "asset": {
                                "properties": {
                                    "assetID": {
                                        "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetStateHistory": {
            "description": "Delete an asset's history from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "asset": {
                                "properties": {
                                    "assetID": {
                                        "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAssetStateHistory"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAsset": {
            "description": "Delete one or more properties from an asset's state",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "asset": {
                                "properties": {
                                    "assetID": {
                                        "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "qprops": {
                                "description": "Qualified property names such as common.location",
                                "items": {
                                    "type": "string"
                                },
                                "type": "array"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deletePropertiesFromAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteWorldState": {
            "description": "**** WARNING *** Clears the entire contents of world state, redeploy the contract after using this, in debugging mode, will require a restart",
            "properties": {
                "args": {
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteWorldState"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "initContract": {
            "description": "Sets contract version and nickname",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "nickname": {
                                "default": "IOT Contract Platform",
                                "description": "The nickname of the current contract instance",
                                "type": "string"
                            },
                            "version": {
                                "description": "The version number of the current contract instance",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "initContract"
                    ],
                    "type": "string"
                },
                "method": "deploy"
            },
            "type": "object"
        },
        "readAllAssets": {
            "description": "Returns the state of all assets, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllAssets"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of asset states, can mix asset classes",
                    "items": {
                        "description": "A asset's complete state",
                        "properties": {
                            "alerts": {
                                "description": "An array of alert names",
                                "items": {
                                    "description": "An alert name",
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "assetID": {
                                "description": "This asset's world state asset ID",
                                "type": "string"
                            },
                            "class": {
                                "description": "An asset's classifier definition",
                                "properties": {
                                    "assetidpath": "An asset's primary key, expressed as a qualified property path (see example contracts)",
                                    "name": "An asset's class name",
                                    "prefix": "An asset's world state prefix, used to allow iteration over all assets of a class"
                                },
                                "type": "object"
                            },
                            "compliant": {
                                "description": "This asset has no active alerts",
                                "type": "boolean"
                            },
                            "eventin": {
                                "description": "The contract event that created this state, for example updateAsset",
                                "properties": {
                                    "asset": {
                                        "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                        "properties": {
                                            "assetID": {
                                                "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "The carrier in possession of this asset",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "Common properties for all assets",
                                                "properties": {
                                                    "appdata": {
                                                        "description": "Application managed information as an array of key:value pairs",
                                                        "items": {
                                                            "properties": {
                                                                "K": {
                                                                    "type": "string"
                                                                },
                                                                "V": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "deviceID": {
                                                        "description": "A unique identifier for the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "devicetimestamp": {
                                                        "description": "A timestamp recoded by the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "location": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of an asset's contents in degrees Celsuis",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "assetID"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "eventout": {
                                "description": "The chaincode event emitted on invoke exit, if any",
                                "properties": {
                                    "asset": {
                                        "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                        "properties": {
                                            "name": {
                                                "default": "EVT.IOTCP.INVOKE.RESULT",
                                                "enum": [
                                                    "EVT.IOTCP.INVOKE.RESULT"
                                                ],
                                                "type": "string"
                                            },
                                            "payload": {
                                                "description": "A map of contributed results",
                                                "properties": {
                                                    "description": "the overall status of the invoke result, defined by err",
                                                    "properties": {
                                                        "activeAlerts": {
                                                            "description": "An array of alert names",
                                                            "items": {
                                                                "description": "An alert name",
                                                                "type": "string"
                                                            },
                                                            "type": "array"
                                                        },
                                                        "alertsCleared": {
                                                            "description": "An array of alert names",
                                                            "items": {
                                                                "description": "An alert name",
                                                                "type": "string"
                                                            },
                                                            "type": "array"
                                                        },
                                                        "alertsRaised": {
                                                            "description": "An array of alert names",
                                                            "items": {
                                                                "description": "An alert name",
                                                                "type": "string"
                                                            },
                                                            "type": "array"
                                                        },
                                                        "invokeresult": {
                                                            "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                            "properties": {
                                                                "message": {
                                                                    "type": "string"
                                                                },
                                                                "status": {
                                                                    "enum": [
                                                                        "OK",
                                                                        "ERROR"
                                                                    ],
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "state": {
                                "description": "Properties that have been received or calculated for this asset",
                                "properties": {
                                    "asset": {
                                        "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                        "properties": {
                                            "assetID": {
                                                "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "The carrier in possession of this asset",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "Common properties for all assets",
                                                "properties": {
                                                    "appdata": {
                                                        "description": "Application managed information as an array of key:value pairs",
                                                        "items": {
                                                            "properties": {
                                                                "K": {
                                                                    "type": "string"
                                                                },
                                                                "V": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "deviceID": {
                                                        "description": "A unique identifier for the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "devicetimestamp": {
                                                        "description": "A timestamp recoded by the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "location": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of an asset's contents in degrees Celsuis",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "assetID"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "txnid": {
                                "description": "Transaction UUID matching the blockchain",
                                "type": "string"
                            },
                            "txnts": {
                                "description": "Transaction timestamp matching the blockchain",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAllRoutes": {
            "description": "Returns an array of registered API calls by function (debugging)",
            "properties": {
                "args": {
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllRoutes"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "An array of routes",
                    "items": {
                        "description": "A route defines a contract API that can be called to perform a service",
                        "properties": {
                            "class": {
                                "description": "An asset's classifier definition",
                                "properties": {
                                    "assetidpath": "An asset's primary key, expressed as a qualified property path (see example contracts)",
                                    "name": "An asset's class name",
                                    "prefix": "An asset's world state prefix, used to allow iteration over all assets of a class"
                                },
                                "type": "object"
                            },
                            "functionname": {
                                "type": "string"
                            },
                            "method": {
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAllRules": {
            "description": "Returns an array of registered rules by class (debugging)",
            "properties": {
                "args": {
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllRules"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "An array of rules",
                    "items": {
                        "description": "A rule defines a behavior that is applied to every new asset state just before writing to world state, often raises or clears alerts",
                        "properties": {
                            "alerts": {
                                "description": "An array of alert names",
                                "items": {
                                    "description": "An alert name",
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "class": {
                                "description": "An asset's classifier definition",
                                "properties": {
                                    "assetidpath": "An asset's primary key, expressed as a qualified property path (see example contracts)",
                                    "name": "An asset's class name",
                                    "prefix": "An asset's world state prefix, used to allow iteration over all assets of a class"
                                },
                                "type": "object"
                            },
                            "rulename": {
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAsset": {
            "description": "Returns the state an asset",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "asset": {
                                "properties": {
                                    "assetID": {
                                        "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAsset"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A asset's complete state",
                    "properties": {
                        "alerts": {
                            "description": "An array of alert names",
                            "items": {
                                "description": "An alert name",
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "assetID": {
                            "description": "This asset's world state asset ID",
                            "type": "string"
                        },
                        "class": {
                            "description": "An asset's classifier definition",
                            "properties": {
                                "assetidpath": "An asset's primary key, expressed as a qualified property path (see example contracts)",
                                "name": "An asset's class name",
                                "prefix": "An asset's world state prefix, used to allow iteration over all assets of a class"
                            },
                            "type": "object"
                        },
                        "compliant": {
                            "description": "This asset has no active alerts",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The contract event that created this state, for example updateAsset",
                            "properties": {
                                "asset": {
                                    "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                    "properties": {
                                        "assetID": {
                                            "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                            "type": "string"
                                        },
                                        "carrier": {
                                            "description": "The carrier in possession of this asset",
                                            "type": "string"
                                        },
                                        "common": {
                                            "description": "Common properties for all assets",
                                            "properties": {
                                                "appdata": {
                                                    "description": "Application managed information as an array of key:value pairs",
                                                    "items": {
                                                        "properties": {
                                                            "K": {
                                                                "type": "string"
                                                            },
                                                            "V": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "deviceID": {
                                                    "description": "A unique identifier for the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "devicetimestamp": {
                                                    "description": "A timestamp recoded by the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "temperature": {
                                            "description": "Temperature of an asset's contents in degrees Celsuis",
                                            "type": "number"
                                        }
                                    },
                                    "required": [
                                        "assetID"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "eventout": {
                            "description": "The chaincode event emitted on invoke exit, if any",
                            "properties": {
                                "asset": {
                                    "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                    "properties": {
                                        "name": {
                                            "default": "EVT.IOTCP.INVOKE.RESULT",
                                            "enum": [
                                                "EVT.IOTCP.INVOKE.RESULT"
                                            ],
                                            "type": "string"
                                        },
                                        "payload": {
                                            "description": "A map of contributed results",
                                            "properties": {
                                                "description": "the overall status of the invoke result, defined by err",
                                                "properties": {
                                                    "activeAlerts": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsCleared": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "alertsRaised": {
                                                        "description": "An array of alert names",
                                                        "items": {
                                                            "description": "An alert name",
                                                            "type": "string"
                                                        },
                                                        "type": "array"
                                                    },
                                                    "invokeresult": {
                                                        "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                        "properties": {
                                                            "message": {
                                                                "type": "string"
                                                            },
                                                            "status": {
                                                                "enum": [
                                                                    "OK",
                                                                    "ERROR"
                                                                ],
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "state": {
                            "description": "Properties that have been received or calculated for this asset",
                            "properties": {
                                "asset": {
                                    "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                    "properties": {
                                        "assetID": {
                                            "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                            "type": "string"
                                        },
                                        "carrier": {
                                            "description": "The carrier in possession of this asset",
                                            "type": "string"
                                        },
                                        "common": {
                                            "description": "Common properties for all assets",
                                            "properties": {
                                                "appdata": {
                                                    "description": "Application managed information as an array of key:value pairs",
                                                    "items": {
                                                        "properties": {
                                                            "K": {
                                                                "type": "string"
                                                            },
                                                            "V": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "deviceID": {
                                                    "description": "A unique identifier for the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "devicetimestamp": {
                                                    "description": "A timestamp recoded by the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "temperature": {
                                            "description": "Temperature of an asset's contents in degrees Celsuis",
                                            "type": "number"
                                        }
                                    },
                                    "required": [
                                        "assetID"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "txnid": {
                            "description": "Transaction UUID matching the blockchain",
                            "type": "string"
                        },
                        "txnts": {
                            "description": "Transaction timestamp matching the blockchain",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetStateHistory": {
            "description": "Returns history for an asset",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "asset": {
                                "properties": {
                                    "assetID": {
                                        "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "daterange": {
                                "description": "if specified, dates must fall in between these values, inclusive",
                                "properties": {
                                    "begin": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    },
                                    "end": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetStateHistory"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of asset states, can mix asset classes",
                    "items": {
                        "description": "A asset's complete state",
                        "properties": {
                            "alerts": {
                                "description": "An array of alert names",
                                "items": {
                                    "description": "An alert name",
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "assetID": {
                                "description": "This asset's world state asset ID",
                                "type": "string"
                            },
                            "class": {
                                "description": "An asset's classifier definition",
                                "properties": {
                                    "assetidpath": "An asset's primary key, expressed as a qualified property path (see example contracts)",
                                    "name": "An asset's class name",
                                    "prefix": "An asset's world state prefix, used to allow iteration over all assets of a class"
                                },
                                "type": "object"
                            },
                            "compliant": {
                                "description": "This asset has no active alerts",
                                "type": "boolean"
                            },
                            "eventin": {
                                "description": "The contract event that created this state, for example updateAsset",
                                "properties": {
                                    "asset": {
                                        "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                        "properties": {
                                            "assetID": {
                                                "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "The carrier in possession of this asset",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "Common properties for all assets",
                                                "properties": {
                                                    "appdata": {
                                                        "description": "Application managed information as an array of key:value pairs",
                                                        "items": {
                                                            "properties": {
                                                                "K": {
                                                                    "type": "string"
                                                                },
                                                                "V": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "deviceID": {
                                                        "description": "A unique identifier for the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "devicetimestamp": {
                                                        "description": "A timestamp recoded by the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "location": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of an asset's contents in degrees Celsuis",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "assetID"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "eventout": {
                                "description": "The chaincode event emitted on invoke exit, if any",
                                "properties": {
                                    "asset": {
                                        "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                        "properties": {
                                            "name": {
                                                "default": "EVT.IOTCP.INVOKE.RESULT",
                                                "enum": [
                                                    "EVT.IOTCP.INVOKE.RESULT"
                                                ],
                                                "type": "string"
                                            },
                                            "payload": {
                                                "description": "A map of contributed results",
                                                "properties": {
                                                    "description": "the overall status of the invoke result, defined by err",
                                                    "properties": {
                                                        "activeAlerts": {
                                                            "description": "An array of alert names",
                                                            "items": {
                                                                "description": "An alert name",
                                                                "type": "string"
                                                            },
                                                            "type": "array"
                                                        },
                                                        "alertsCleared": {
                                                            "description": "An array of alert names",
                                                            "items": {
                                                                "description": "An alert name",
                                                                "type": "string"
                                                            },
                                                            "type": "array"
                                                        },
                                                        "alertsRaised": {
                                                            "description": "An array of alert names",
                                                            "items": {
                                                                "description": "An alert name",
                                                                "type": "string"
                                                            },
                                                            "type": "array"
                                                        },
                                                        "invokeresult": {
                                                            "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                            "properties": {
                                                                "message": {
                                                                    "type": "string"
                                                                },
                                                                "status": {
                                                                    "enum": [
                                                                        "OK",
                                                                        "ERROR"
                                                                    ],
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "state": {
                                "description": "Properties that have been received or calculated for this asset",
                                "properties": {
                                    "asset": {
                                        "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                        "properties": {
                                            "assetID": {
                                                "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "The carrier in possession of this asset",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "Common properties for all assets",
                                                "properties": {
                                                    "appdata": {
                                                        "description": "Application managed information as an array of key:value pairs",
                                                        "items": {
                                                            "properties": {
                                                                "K": {
                                                                    "type": "string"
                                                                },
                                                                "V": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "deviceID": {
                                                        "description": "A unique identifier for the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "devicetimestamp": {
                                                        "description": "A timestamp recoded by the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "location": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of an asset's contents in degrees Celsuis",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "assetID"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "txnid": {
                                "description": "Transaction UUID matching the blockchain",
                                "type": "string"
                            },
                            "txnts": {
                                "description": "Transaction timestamp matching the blockchain",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readRecentStates": {
            "description": "Returns the state of recently updated assets",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "begin": {
                                "description": "zero based beginning of range",
                                "type": "integer"
                            },
                            "end": {
                                "description": "zero based end of range, absence means to end",
                                "type": "integer"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readRecentStates"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of asset states, can mix asset classes",
                    "items": {
                        "description": "A asset's complete state",
                        "properties": {
                            "alerts": {
                                "description": "An array of alert names",
                                "items": {
                                    "description": "An alert name",
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "assetID": {
                                "description": "This asset's world state asset ID",
                                "type": "string"
                            },
                            "class": {
                                "description": "An asset's classifier definition",
                                "properties": {
                                    "assetidpath": "An asset's primary key, expressed as a qualified property path (see example contracts)",
                                    "name": "An asset's class name",
                                    "prefix": "An asset's world state prefix, used to allow iteration over all assets of a class"
                                },
                                "type": "object"
                            },
                            "compliant": {
                                "description": "This asset has no active alerts",
                                "type": "boolean"
                            },
                            "eventin": {
                                "description": "The contract event that created this state, for example updateAsset",
                                "properties": {
                                    "asset": {
                                        "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                        "properties": {
                                            "assetID": {
                                                "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "The carrier in possession of this asset",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "Common properties for all assets",
                                                "properties": {
                                                    "appdata": {
                                                        "description": "Application managed information as an array of key:value pairs",
                                                        "items": {
                                                            "properties": {
                                                                "K": {
                                                                    "type": "string"
                                                                },
                                                                "V": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "deviceID": {
                                                        "description": "A unique identifier for the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "devicetimestamp": {
                                                        "description": "A timestamp recoded by the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "location": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of an asset's contents in degrees Celsuis",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "assetID"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "eventout": {
                                "description": "The chaincode event emitted on invoke exit, if any",
                                "properties": {
                                    "asset": {
                                        "description": "A chaincode event defining the standard platform-generated result event for a contract invoke, contains an array of contributed results",
                                        "properties": {
                                            "name": {
                                                "default": "EVT.IOTCP.INVOKE.RESULT",
                                                "enum": [
                                                    "EVT.IOTCP.INVOKE.RESULT"
                                                ],
                                                "type": "string"
                                            },
                                            "payload": {
                                                "description": "A map of contributed results",
                                                "properties": {
                                                    "description": "the overall status of the invoke result, defined by err",
                                                    "properties": {
                                                        "activeAlerts": {
                                                            "description": "An array of alert names",
                                                            "items": {
                                                                "description": "An alert name",
                                                                "type": "string"
                                                            },
                                                            "type": "array"
                                                        },
                                                        "alertsCleared": {
                                                            "description": "An array of alert names",
                                                            "items": {
                                                                "description": "An alert name",
                                                                "type": "string"
                                                            },
                                                            "type": "array"
                                                        },
                                                        "alertsRaised": {
                                                            "description": "An array of alert names",
                                                            "items": {
                                                                "description": "An alert name",
                                                                "type": "string"
                                                            },
                                                            "type": "array"
                                                        },
                                                        "invokeresult": {
                                                            "description": "status: OK==txn succeeded, ERROR==txn failed",
                                                            "properties": {
                                                                "message": {
                                                                    "type": "string"
                                                                },
                                                                "status": {
                                                                    "enum": [
                                                                        "OK",
                                                                        "ERROR"
                                                                    ],
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "state": {
                                "description": "Properties that have been received or calculated for this asset",
                                "properties": {
                                    "asset": {
                                        "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                        "properties": {
                                            "assetID": {
                                                "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "The carrier in possession of this asset",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "Common properties for all assets",
                                                "properties": {
                                                    "appdata": {
                                                        "description": "Application managed information as an array of key:value pairs",
                                                        "items": {
                                                            "properties": {
                                                                "K": {
                                                                    "type": "string"
                                                                },
                                                                "V": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "deviceID": {
                                                        "description": "A unique identifier for the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "devicetimestamp": {
                                                        "description": "A timestamp recoded by the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "location": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of an asset's contents in degrees Celsuis",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "assetID"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "txnid": {
                                "description": "Transaction UUID matching the blockchain",
                                "type": "string"
                            },
                            "txnts": {
                                "description": "Transaction timestamp matching the blockchain",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readWorldState": {
            "description": "Returns the entire contents of world state",
            "properties": {
                "args": {
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readWorldState"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "properties": {},
                    "type": "object"
                }
            },
            "type": "object"
        },
        "replaceAsset": {
            "description": "Replaces an asset's state (e.g. put existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "asset": {
                                "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                "properties": {
                                    "assetID": {
                                        "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                        "type": "string"
                                    },
                                    "carrier": {
                                        "description": "The carrier in possession of this asset",
                                        "type": "string"
                                    },
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
                                            },
                                            "location": {
                                                "description": "A geographical coordinate",
                                                "properties": {
                                                    "latitude": {
                                                        "type": "number"
                                                    },
                                                    "longitude": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "temperature": {
                                        "description": "Temperature of an asset's contents in degrees Celsuis",
                                        "type": "number"
                                    }
                                },
                                "required": [
                                    "assetID"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "replaceAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "setCreateOnFirstUpdate": {
            "description": "Allow updateAsset to create an asset upon receipt of its first event",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "setCreateOnFirstUpdate": {
                                "description": "Allows updates to create missing assets on first event",
                                "type": "boolean"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "setCreateOnFirstUpdate"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "setLoggingLevel": {
            "description": "Sets the logging level for the contract",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "logLevel": {
                                "enum": [
                                    "CRITICAL",
                                    "ERROR",
                                    "WARNING",
                                    "NOTICE",
                                    "INFO",
                                    "DEBUG"
                                ],
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "setLoggingLevel"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateAsset": {
            "description": "Update an asset's state with one or more property changes",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "asset": {
                                "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                "properties": {
                                    "assetID": {
                                        "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                        "type": "string"
                                    },
                                    "carrier": {
                                        "description": "The carrier in possession of this asset",
                                        "type": "string"
                                    },
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
                                            },
                                            "location": {
                                                "description": "A geographical coordinate",
                                                "properties": {
                                                    "latitude": {
                                                        "type": "number"
                                                    },
                                                    "longitude": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "temperature": {
                                        "description": "Temperature of an asset's contents in degrees Celsuis",
                                        "type": "number"
                                    }
                                },
                                "required": [
                                    "assetID"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "updateAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        }
    },
    "Model": {
        "asset": {
            "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
            "properties": {
                "assetID": {
                    "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                    "type": "string"
                },
                "carrier": {
                    "description": "The carrier in possession of this asset",
                    "type": "string"
                },
                "common": {
                    "description": "Common properties for all assets",
                    "properties": {
                        "appdata": {
                            "description": "Application managed information as an array of key:value pairs",
                            "items": {
                                "properties": {
                                    "K": {
                                        "type": "string"
                                    },
                                    "V": {
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "deviceID": {
                            "description": "A unique identifier for the device that sent the current event",
                            "type": "string"
                        },
                        "devicetimestamp": {
                            "description": "A timestamp recoded by the device that sent the current event",
                            "type": "string"
                        },
                        "location": {
                            "description": "A geographical coordinate",
                            "properties": {
                                "latitude": {
                                    "type": "number"
                                },
                                "longitude": {
                                    "type": "number"
                                }
                            },
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "temperature": {
                    "description": "Temperature of an asset's contents in degrees Celsuis",
                    "type": "number"
                }
            },
            "required": [
                "assetID"
            ],
            "type": "object"
        }
    }
}`


	var readAssetSchemas iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		return []byte(schemas), nil
	}
	func init() {
		iot.AddRoute("readAssetSchemas", "query", iot.SystemClass, readAssetSchemas)
	}
	