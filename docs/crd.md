# API Docs
This Document documents the types introduced by the Fluent Bit Operator to be consumed by users.
> Note this document is generated from code comments. When contributing a change to this document please do so by changing the code comments.
## Table of Contents
* [Decorder](#decorder)
* [Filter](#filter)
* [FilterItem](#filteritem)
* [FilterList](#filterlist)
* [FilterSpec](#filterspec)
* [FluentBit](#fluentbit)
* [FluentBitConfig](#fluentbitconfig)
* [FluentBitConfigList](#fluentbitconfiglist)
* [FluentBitConfigSpec](#fluentbitconfigspec)
* [FluentBitList](#fluentbitlist)
* [FluentBitSpec](#fluentbitspec)
* [Input](#input)
* [InputList](#inputlist)
* [InputSpec](#inputspec)
* [Output](#output)
* [OutputList](#outputlist)
* [OutputSpec](#outputspec)
* [Parser](#parser)
* [ParserList](#parserlist)
* [ParserSpec](#parserspec)
* [Service](#service)
## Decorder




| Field | Description | Scheme |
| ----- | ----------- | ------ |
| decodeField | If the content can be decoded in a structured message, append that structure message (keys and values) to the original log message. | string |
| decodeFieldAs | Any content decoded (unstructured or structured) will be replaced in the same key/value, no extra keys are added. | string |

[Back to TOC](#table-of-contents)
## Filter

Filter defines a Filter configuration.


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| metadata |  | [metav1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#objectmeta-v1-meta) |
| spec | Specification of desired Filter configuration. | FilterSpec |

[Back to TOC](#table-of-contents)
## FilterItem




| Field | Description | Scheme |
| ----- | ----------- | ------ |
| grep | Grep defines Grep Filter configuration. | *[filter.Grep](plugins/fluentbit/filter/grep.md) |
| recordModifier | RecordModifier defines Record Modifier Filter configuration. | *[filter.RecordModifier](plugins/fluentbit/filter/recordmodifier.md) |
| kubernetes | Kubernetes defines Kubernetes Filter configuration. | *[filter.Kubernetes](plugins/fluentbit/filter/kubernetes.md) |
| modify | Modify defines Modify Filter configuration. | *[filter.Modify](plugins/fluentbit/filter/modify.md) |
| nest | Nest defines Nest Filter configuration. | *[filter.Nest](plugins/fluentbit/filter/nest.md) |
| parser | Parser defines Parser Filter configuration. | *[filter.Parser](plugins/fluentbit/filter/parser.md) |
| lua | Lua defines Lua Filter configuration. | *[filter.Lua](plugins/fluentbit/filter/lua.md) |

[Back to TOC](#table-of-contents)
## FilterList

FilterList contains a list of Filter


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| metadata |  | [metav1.ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#listmeta-v1-meta) |
| items |  | []Filter |

[Back to TOC](#table-of-contents)
## FilterSpec

FilterSpec defines the desired state of Filter


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| match | A pattern to match against the tags of incoming records. It's case sensitive and support the star (*) character as a wildcard. | string |
| matchRegex | A regular expression to match against the tags of incoming records. Use this option if you want to use the full regex syntax. | string |
| filters | A set of filter plugins in order. | []FilterItem |

[Back to TOC](#table-of-contents)
## FluentBit

FluentBit is the Schema for the fluentbits API


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| metadata |  | [metav1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#objectmeta-v1-meta) |
| spec |  | FluentBitSpec |
| status |  | FluentBitStatus |

[Back to TOC](#table-of-contents)
## FluentBitConfig

FluentBitConfig is the Schema for the fluentbitconfigs API


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| metadata |  | [metav1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#objectmeta-v1-meta) |
| spec |  | FluentBitConfigSpec |

[Back to TOC](#table-of-contents)
## FluentBitConfigList

FluentBitConfigList contains a list of FluentBitConfig


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| metadata |  | [metav1.ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#listmeta-v1-meta) |
| items |  | []FluentBitConfig |

[Back to TOC](#table-of-contents)
## FluentBitConfigSpec

FluentBitConfigSpec defines the desired state of FluentBitConfig


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| service | Service defines the global behaviour of the Fluent Bit engine. | *Service |
| inputSelector | Select input plugins | [metav1.LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#labelselector-v1-meta) |
| filterSelector | Select filter plugins | [metav1.LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#labelselector-v1-meta) |
| outputSelector | Select output plugins | [metav1.LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#labelselector-v1-meta) |
| parserSelector | Select parser plugins | [metav1.LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#labelselector-v1-meta) |

[Back to TOC](#table-of-contents)
## FluentBitList

FluentBitList contains a list of FluentBit


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| metadata |  | [metav1.ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#listmeta-v1-meta) |
| items |  | []FluentBit |

[Back to TOC](#table-of-contents)
## FluentBitSpec

FluentBitSpec defines the desired state of FluentBit


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| image | Fluent Bit image. | string |
| imagePullPolicy | Fluent Bit image pull policy. | corev1.PullPolicy |
| positionDB | Storage for position db. You will use it if tail input is enabled. | [corev1.VolumeSource](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#volume-v1-core) |
| containerLogRealPath | Container log path | string |
| nodeSelector | NodeSelector | map[string]string |
| tolerations | Tolerations | [][corev1.Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#toleration-v1-core) |
| fluentBitConfigName | Fluentbitconfig object associated with this Fluentbit | string |
| secrets | The Secrets are mounted into /fluent-bit/secrets/<secret-name>. | []string |
| runtimeClassName | The runtimeClassName represents the container runtime configuration. | string |
| priorityClassName | The priorityClassName represents the pod's priority. | string |

[Back to TOC](#table-of-contents)
## Input

Input is the Schema for the inputs API


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| metadata |  | [metav1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#objectmeta-v1-meta) |
| spec |  | InputSpec |

[Back to TOC](#table-of-contents)
## InputList

InputList contains a list of Input


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| metadata |  | [metav1.ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#listmeta-v1-meta) |
| items |  | []Input |

[Back to TOC](#table-of-contents)
## InputSpec

InputSpec defines the desired state of Input


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| alias | A user friendly alias name for this input plugin. | string |
| dummy | Dummy defines Dummy Input configuration. | *[input.Dummy](plugins/fluentbit/input/dummy.md) |
| tail | Tail defines Tail Input configuration. | *[input.Tail](plugins/fluentbit/input/tail.md) |
| systemd | Systemd defines Systemd Input configuration. | *[input.Systemd](plugins/fluentbit/input/systemd.md) |

[Back to TOC](#table-of-contents)
## Output

Output is the Schema for the outputs API


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| metadata |  | [metav1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#objectmeta-v1-meta) |
| spec |  | OutputSpec |

[Back to TOC](#table-of-contents)
## OutputList

OutputList contains a list of Output


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| metadata |  | [metav1.ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#listmeta-v1-meta) |
| items |  | []Output |

[Back to TOC](#table-of-contents)
## OutputSpec

OutputSpec defines the desired state of Output


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| match | A pattern to match against the tags of incoming records. It's case sensitive and support the star (*) character as a wildcard. | string |
| matchRegex | A regular expression to match against the tags of incoming records. Use this option if you want to use the full regex syntax. | string |
| alias | A user friendly alias name for this output plugin. | string |
| retry_limit | This option allows to disable retries or impose a limit to try N times and then discard the data after reaching that limit. | string |
| es | Elasticsearch defines Elasticsearch Output configuration. | *[output.Elasticsearch](plugins/fluentbit/output/elasticsearch.md) |
| file | File defines File Output configuration. | *[output.File](plugins/fluentbit/output/file.md) |
| forward | Forward defines Forward Output configuration. | *[output.Forward](plugins/fluentbit/output/forward.md) |
| http | HTTP defines HTTP Output configuration. | *[output.HTTP](plugins/fluentbit/output/http.md) |
| kafka | Kafka defines Kafka Output configuration. | *[output.Kafka](plugins/fluentbit/output/kafka.md) |
| null | Null defines Null Output configuration. | *[output.Null](plugins/fluentbit/output/null.md) |
| stdout | Stdout defines Stdout Output configuration. | *[output.Stdout](plugins/fluentbit/output/stdout.md) |
| tcp | TCP defines TCP Output configuration. | *[output.TCP](plugins/fluentbit/output/tcp.md) |
| loki | Loki defines Loki Output configuration. | *[output.Loki](plugins/fluentbit/output/loki.md) |
| syslog | Syslog defines Syslog Output configuration. | *[output.Syslog](plugins/fluentbit/output/syslog.md) |
| datadog | DataDog defines DataDog Output configuration. | *[output.DataDog](plugins/fluentbit/output/datadog.md) |

[Back to TOC](#table-of-contents)
## Parser

Parser is the Schema for the parsers API


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| metadata |  | [metav1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#objectmeta-v1-meta) |
| spec |  | ParserSpec |

[Back to TOC](#table-of-contents)
## ParserList

ParserList contains a list of Parser


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| metadata |  | [metav1.ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#listmeta-v1-meta) |
| items |  | []Parser |

[Back to TOC](#table-of-contents)
## ParserSpec

ParserSpec defines the desired state of Parser


| Field | Description | Scheme |
| ----- | ----------- | ------ |
| json | JSON defines json parser configuration. | *[parser.JSON](plugins/fluentbit/parser/json.md) |
| regex | Regex defines regex parser configuration. | *[parser.Regex](plugins/fluentbit/parser/regex.md) |
| ltsv | LTSV defines ltsv parser configuration. | *[parser.LSTV](plugins/fluentbit/parser/lstv.md) |
| logfmt | Logfmt defines logfmt parser configuration. | *[parser.Logfmt](plugins/fluentbit/parser/logfmt.md) |
| decoders | Decoders are a built-in feature available through the Parsers file, each Parser definition can optionally set one or multiple decoders. There are two type of decoders type: Decode_Field and Decode_Field_As. | []Decorder |

[Back to TOC](#table-of-contents)
## Service




| Field | Description | Scheme |
| ----- | ----------- | ------ |
| daemon | If true go to background on start | *bool |
| flushSeconds | Interval to flush output | *int64 |
| graceSeconds | Wait time on exit | *int64 |
| httpListen | Address to listen | string |
| httpPort | Port to listen | *int32 |
| httpServer | If true enable statistics HTTP server | *bool |
| logFile | File to log diagnostic output | string |
| logLevel | Diagnostic level (error/warning/info/debug/trace) | string |
| parsersFile | Optional 'parsers' config file (can be multiple) | string |

[Back to TOC](#table-of-contents)
