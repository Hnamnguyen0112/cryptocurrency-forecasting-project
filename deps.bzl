load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_actgardner_gogen_avro_v10",
        importpath = "github.com/actgardner/gogen-avro/v10",
        sum = "h1:z3pOGblRjAJCYpkIJ8CmbMJdksi4rAhaygw0dyXZ930=",
        version = "v10.2.1",
    )
    go_repository(
        name = "com_github_azure_go_ansiterm",
        importpath = "github.com/Azure/go-ansiterm",
        sum = "h1:UQHMgLO+TxOElx5B5HZ4hJQsoJ/PvUvKRhJHDQXO8P8=",
        version = "v0.0.0-20210617225240-d185dfc1b5a1",
    )
    go_repository(
        name = "com_github_bahlo_generic_list_go",
        importpath = "github.com/bahlo/generic-list-go",
        sum = "h1:5sz/EEAK+ls5wF+NeqDpk5+iNdMDXrh3z3nPnH1Wvgk=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_buger_jsonparser",
        importpath = "github.com/buger/jsonparser",
        sum = "h1:2PnMjfWD7wBILjqQbt530v576A/cAbQvEW9gGIpYMUs=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v4",
        importpath = "github.com/cenkalti/backoff/v4",
        sum = "h1:cFAlzYUlVYDysBEH2T5hyJZMh3+5+WCBvSnK6Q8UtC4=",
        version = "v4.1.3",
    )
    go_repository(
        name = "com_github_confluentinc_confluent_kafka_go_v2",
        importpath = "github.com/confluentinc/confluent-kafka-go/v2",
        patches = ["//bazel:kafka.patch"],
        sum = "h1:icCHutJouWlQREayFwCc7lxDAhws08td+W3/gdqgZts=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_github_containerd_cgroups",
        importpath = "github.com/containerd/cgroups",
        sum = "h1:jN/mbWBEaz+T1pi5OFtnkQ+8qnmEbAr1Oo1FRm5B0dA=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_containerd_containerd",
        importpath = "github.com/containerd/containerd",
        sum = "h1:h4dOFDwzHmqFEP754PgfgTeVXFnLiRc6kiqC7tplDJs=",
        version = "v1.6.8",
    )
    go_repository(
        name = "com_github_creack_pty",
        importpath = "github.com/creack/pty",
        sum = "h1:uDmaGzcdjhF4i/plgjmEsriH11Y0o7RKapEf/LDaM3w=",
        version = "v1.1.9",
    )
    go_repository(
        name = "com_github_data_dog_go_sqlmock",
        importpath = "github.com/DATA-DOG/go-sqlmock",
        sum = "h1:OcvFkGmslmlZibjAjaHm3L//6LiuBgolP7OputlJIzU=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_docker_distribution",
        importpath = "github.com/docker/distribution",
        sum = "h1:Q50tZOPR6T/hjNsyc9g8/syEs6bk8XXApsHjKukMl68=",
        version = "v2.8.1+incompatible",
    )
    go_repository(
        name = "com_github_docker_docker",
        importpath = "github.com/docker/docker",
        sum = "h1:JYCuMrWaVNophQTOrMMoSwudOVEfcegoZZrleKc1xwE=",
        version = "v20.10.17+incompatible",
    )
    go_repository(
        name = "com_github_docker_go_connections",
        importpath = "github.com/docker/go-connections",
        sum = "h1:El9xVISelRB7BuFusrZozjnkIM5YnzCViNKohAFqRJQ=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_docker_go_units",
        importpath = "github.com/docker/go-units",
        sum = "h1:69rxXcBk27SvSaaxTtLh/8llcHD8vYHT7WSdRZ/jvr4=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_golang_groupcache",
        importpath = "github.com/golang/groupcache",
        sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
        version = "v0.0.0-20210331224755-41bb18bfe9da",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:KhyjKVUg7Usr/dYsdSqoFveMYd5ko72D+zANwlG1mmg=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_gorilla_websocket",
        importpath = "github.com/gorilla/websocket",
        sum = "h1:saDtZ6Pbx/0u+bgYQ3q96pZgCzfhKXGPqt7kZ72aNNg=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_github_heetch_avro",
        importpath = "github.com/heetch/avro",
        sum = "h1:5PmgDy1cX/MegMy6btJ4bUFHgT5GLfSYfc5U7+JUQzg=",
        version = "v0.4.4",
    )
    go_repository(
        name = "com_github_iancoleman_orderedmap",
        importpath = "github.com/iancoleman/orderedmap",
        sum = "h1:i462o439ZjprVSFSZLZxcsoAe592sZB1rci2Z8j4wdk=",
        version = "v0.0.0-20190318233801-ac98e3ecb4b0",
    )
    go_repository(
        name = "com_github_invopop_jsonschema",
        importpath = "github.com/invopop/jsonschema",
        sum = "h1:6ovsNSuvn9wEQVOyc72aycBMVQFKz7cPdMJn10CvzRI=",
        version = "v0.12.0",
    )
    go_repository(
        name = "com_github_jackc_pgpassfile",
        importpath = "github.com/jackc/pgpassfile",
        sum = "h1:/6Hmqy13Ss2zCq62VdNG8tM1wchn8zjSGOBJ6icpsIM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_pgservicefile",
        importpath = "github.com/jackc/pgservicefile",
        sum = "h1:bbPeKD0xmW/Y25WS6cokEszi5g+S0QxI/d45PkRi7Nk=",
        version = "v0.0.0-20221227161230-091c0ba34f0a",
    )
    go_repository(
        name = "com_github_jackc_pgx_v5",
        importpath = "github.com/jackc/pgx/v5",
        sum = "h1:amBjrZVmksIdNjxGW/IiIMzxMKZFelXbUoPNb+8sjQw=",
        version = "v5.5.5",
    )
    go_repository(
        name = "com_github_jackc_puddle_v2",
        importpath = "github.com/jackc/puddle/v2",
        sum = "h1:RhxXJtFG022u4ibrCSMSiu5aOq1i77R3OHKNJj77OAk=",
        version = "v2.2.1",
    )
    go_repository(
        name = "com_github_jhump_protoreflect",
        importpath = "github.com/jhump/protoreflect",
        sum = "h1:N88q7JkxTHWFEqReuTsYH1dPIwXxA0ITNQp7avLY10s=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_github_jinzhu_inflection",
        importpath = "github.com/jinzhu/inflection",
        sum = "h1:K317FqzuhWc8YvSVlFMCCUb36O/S9MCKRDI7QkRKD/E=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jinzhu_now",
        importpath = "github.com/jinzhu/now",
        sum = "h1:/o9tlHleP7gOFmsnYNz3RGnqzefHA47wQpKrrdTIwXQ=",
        version = "v1.1.5",
    )
    go_repository(
        name = "com_github_joho_godotenv",
        importpath = "github.com/joho/godotenv",
        sum = "h1:7eLL/+HRGLY0ldzfGMeQkb7vMd0as4CfYvUVzLqw0N0=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_josharian_intern",
        importpath = "github.com/josharian/intern",
        sum = "h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_kisielk_sqlstruct",
        importpath = "github.com/kisielk/sqlstruct",
        sum = "h1:veS9QfglfvqAw2e+eeNT/SbGySq8ajECXJ9e4fPoLhY=",
        version = "v0.0.0-20201105191214-5f3e10d3ab46",
    )
    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        sum = "h1:WgNl7dwNpEZ6jJ9k1snq4pZsg7DOEN8hP9Xw0Tsjwk0=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_magiconair_properties",
        importpath = "github.com/magiconair/properties",
        sum = "h1:5ibWZ6iY0NctNGWo87LalDlEZ6R41TqbbDamhfG/Qzo=",
        version = "v1.8.6",
    )
    go_repository(
        name = "com_github_mailru_easyjson",
        importpath = "github.com/mailru/easyjson",
        sum = "h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=",
        version = "v0.7.7",
    )
    go_repository(
        name = "com_github_microsoft_go_winio",
        importpath = "github.com/Microsoft/go-winio",
        sum = "h1:a9IhgEQBCUEk6QCdml9CiJGhAws+YwffDHEMp1VMrpA=",
        version = "v0.5.2",
    )
    go_repository(
        name = "com_github_microsoft_hcsshim",
        importpath = "github.com/Microsoft/hcsshim",
        sum = "h1:mnUj0ivWy6UzbB1uLFqKR6F+ZyiDc7j4iGgHTpO+5+I=",
        version = "v0.9.4",
    )
    go_repository(
        name = "com_github_moby_sys_mount",
        importpath = "github.com/moby/sys/mount",
        sum = "h1:fX1SVkXFJ47XWDoeFW4Sq7PdQJnV2QIDZAqjNqgEjUs=",
        version = "v0.3.3",
    )
    go_repository(
        name = "com_github_moby_sys_mountinfo",
        importpath = "github.com/moby/sys/mountinfo",
        sum = "h1:BzJjoreD5BMFNmD9Rus6gdd1pLuecOFPt8wC+Vygl78=",
        version = "v0.6.2",
    )
    go_repository(
        name = "com_github_moby_term",
        importpath = "github.com/moby/term",
        sum = "h1:dcztxKSvZ4Id8iPpHERQBbIJfabdt4wUm5qy3wOL2Zc=",
        version = "v0.0.0-20210619224110-3f7ff695adc6",
    )
    go_repository(
        name = "com_github_morikuni_aec",
        importpath = "github.com/morikuni/aec",
        sum = "h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_opencontainers_go_digest",
        importpath = "github.com/opencontainers/go-digest",
        sum = "h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_opencontainers_image_spec",
        importpath = "github.com/opencontainers/image-spec",
        sum = "h1:rc3tiVYb5z54aKaDfakKn0dDjIyPpTtszkjuMzyt7ec=",
        version = "v1.0.3-0.20211202183452-c5a74bcca799",
    )
    go_repository(
        name = "com_github_opencontainers_runc",
        importpath = "github.com/opencontainers/runc",
        sum = "h1:vIXrkId+0/J2Ymu2m7VjGvbSlAId9XNRPhn2p4b+d8w=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_rogpeppe_go_internal",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:exVL4IDcn6na9z1rAb56Vxr+CgyK3nn3O+epU5NdKM8=",
        version = "v1.12.0",
    )
    go_repository(
        name = "com_github_santhosh_tekuri_jsonschema_v5",
        importpath = "github.com/santhosh-tekuri/jsonschema/v5",
        sum = "h1:uIkTLo0AGRc8l7h5l9r+GcYi9qfVPt6lD4/bhmzfiKo=",
        version = "v5.3.0",
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:dJKuHgqk1NNQlqoA6BTlM1Wf9DOH3NBjQyu0h9+AZZE=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:+h33VjcLVPDHtOdpUCuF+7gSuG3yGIftsP1YvFihtJ8=",
        version = "v1.8.2",
    )
    go_repository(
        name = "com_github_testcontainers_testcontainers_go",
        importpath = "github.com/testcontainers/testcontainers-go",
        sum = "h1:h0D5GaYG9mhOWr2qHdEKDXpkce/VlvaYOCzTRi6UBi8=",
        version = "v0.14.0",
    )
    go_repository(
        name = "com_github_wk8_go_ordered_map_v2",
        importpath = "github.com/wk8/go-ordered-map/v2",
        sum = "h1:5h/BUHu93oj4gIdvHHHGsScSTMijfx5PeYkE/fJgbpc=",
        version = "v2.1.8",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
        version = "v1.0.0-20201130134442-10cb98267c6c",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "io_gorm_driver_postgres",
        importpath = "gorm.io/driver/postgres",
        sum = "h1:DkegyItji119OlcaLjqN11kHoUgZ/j13E0jkJZgD6A8=",
        version = "v1.5.9",
    )
    go_repository(
        name = "io_gorm_gorm",
        importpath = "gorm.io/gorm",
        sum = "h1:/Wfyg1B/je1hnDx3sMkX+gAlxrlZpn6X0BXRlwXlvHg=",
        version = "v1.25.11",
    )
    go_repository(
        name = "io_opencensus_go",
        importpath = "go.opencensus.io",
        sum = "h1:gqCw0LfLxScz8irSi8exQc7fyQ0fKQU/qnC/X8+V/1M=",
        version = "v0.23.0",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:0BOZf6qNozI3pkN3fJLwNubheHJYHhMh91GRFOWWK08=",
        version = "v0.0.0-20230331144136-dcfb400f0633",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:EhTqbhiYeixwWQtAEZAxmV9MGqcjEU2mFx52xCzNyag=",
        version = "v1.54.0",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        sum = "h1:kPPoIgf3TsEvrm0PFe15JQ+570QVxYzEvvHqChK+cng=",
        version = "v1.30.0",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:g1v0xeRhjcugydODzvb3mEM9SQ0HGp9s/nh3COQ/C30=",
        version = "v0.22.0",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:KENHtAZL2y3NLMYZeHY9DW8HW8V+kQyJsY/V9JlKvCs=",
        version = "v0.9.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:AQyQV4dYCvJ7vGmJyKki9+PBdyvhkSd8EIx/qb0AYv4=",
        version = "v0.21.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:YsImfSBoP9QPYL0xyKJPq0gcaJdG3rInoqxTWbfQu9M=",
        version = "v0.7.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:q5f1RH2jigJ1MoAWp2KTp3gm5zAGFUTarQZ5U386+4o=",
        version = "v0.19.0",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:+ThwsDv+tYfnJFhF4L8jITxu1tdTWRTZpdsWgEgjL6Q=",
        version = "v0.19.0",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:ScX5w1eTa3QqT8oi6+ziP7dTV1S2+ALU0bI+0zXKWiQ=",
        version = "v0.14.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        sum = "h1:BOw41kyTf3PuCW1pVQf8+Cyg8pMlkYB1oo9iJ6D/lKM=",
        version = "v0.6.0",
    )
