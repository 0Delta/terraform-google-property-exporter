# terraform-google-property-exporter

hashicorp/terraform-provider-google のプロパティを列挙するプログラム

# Using

```
git clone https://github.com/0Delta/terraform-google-property-exporter.git
git switch export-${version}
go run ./
```

## Options

Usage of terraform-google-property-exporter:  
  -csv  
        output by csv (overwrite separator option)  
  -separator string  
        separator char (default "\t")  

## Output Sample

最初の列は種別を表します。それぞれ以下の意味を持ちます。  
+ R - Require  
  必須オプションです。ただし、上位のオプションが非必須であればその単位で省略できます。  
+ O - Optional  
  非必須オプションです。指定しなかった場合は空の値が設定されます。  
+ C - Compute  
  動的に設定される値です。読取専用の値などです。  
+ CO - Compute Optional  
  非必須オプションですが、指定しなかった場合は自動的にデフォルト値が設定される値です。  

```
O       google_compute_global_address   network
C       google_compute_global_address   creation_timestamp
C       google_compute_global_address   self_link
R       google_compute_global_address   name
CO      google_compute_global_address   address
O       google_compute_global_address   ip_version
O       google_compute_global_address   purpose
CO      google_compute_global_address   project
O       google_compute_global_address   address_type
O       google_compute_global_address   description
```

if with `-csv` flag
```
R,google_bigquery_table_iam_member,member
O,google_bigquery_table_iam_member,condition
R,google_bigquery_table_iam_member,condition,expression
R,google_bigquery_table_iam_member,condition,title
O,google_bigquery_table_iam_member,condition,description
C,google_bigquery_table_iam_member,etag
CO,google_bigquery_table_iam_member,project
R,google_bigquery_table_iam_member,dataset_id
R,google_bigquery_table_iam_member,table_id
R,google_bigquery_table_iam_member,role
```

## Core branches (for developer)

+ main
+ patch
+ base

