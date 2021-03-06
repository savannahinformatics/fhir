{
  "resourceType" : "CapabilityStatement",
  "url" : "/metadata",
  "version" : "1",
  "name" : "GoFHIR Capability Statement",
  "status" : "active",
  "date" : "<Date>",
  "publisher" : "The MITRE Corporation",
  "description" : "GoFHIR capability statement",
  "kind" : "instance",
  "fhirVersion" : "1.8.0",
  "acceptUnknown" : "extensions",
  "format" : [
    "json"
  ],
  "rest" : [
    {
      "mode" : "server",
      "resource" : [

        <trunc(ResourceSearchInfos): {r | 
        <\u007B>
          "type" : "<r.name>",
          "interaction" : [
            <\u007B>"code" : "create"<\u007D>,
            <\u007B>"code" : "read"<\u007D>,
            <\u007B>"code" : "update"<\u007D>,
            <\u007B>"code" : "delete"<\u007D>
          ],
          "versioning" : "no-version",
          "conditionalCreate" : true,
          "conditionalUpdate" : true,
          "conditionalDelete" : "multiple",
          "searchInclude" : [

            <r.searchParams: {s |
            <if(s.targets)><! we use the existence of s.targets to infer a "reference" parameter type !>
            "<r.name>:<s.name>",
            <endif>
            }>
            "*"
          ],
          "transactionMode" : "both",
          "searchParam" : [

            <trunc(r.searchParams): {s |
            <\u007B>
              "name" : "<s.name>",
              "type" : "<s.type>"<if(s.targets)>,
              "target" : [
                <trunc(s.targets): {t |
                "<t>",
                }>
                <last(s.targets): {t |
                "<t>"
                }>
              ]
              <endif>
            <\u007D>,
            }>

            <last(r.searchParams): {s |
            <\u007B>
              "name" : "<s.name>",
              "type" : "<s.type>"<if(s.targets)>,
              "target" : [
                <trunc(s.targets): {t |
                "<t>",
                }>
                <last(s.targets): {t |
                "<t>"
                }>
              ]
              <endif>
            <\u007D>
            }>
          ]
        <\u007D>,
        }>

        <last(ResourceSearchInfos): {r | 
        <\u007B>
          "type" : "<r.name>",
          "interaction" : [
            <\u007B>"code" : "create"<\u007D>,
            <\u007B>"code" : "read"<\u007D>,
            <\u007B>"code" : "update"<\u007D>,
            <\u007B>"code" : "delete"<\u007D>
          ],
          "versioning" : "no-version",
          "conditionalCreate" : true,
          "conditionalUpdate" : true,
          "conditionalDelete" : "multiple",
          "searchInclude" : [

            <r.searchParams: {s |
            <if(s.targets)><! we use the existence of s.targets to infer a "reference" parameter type !>
            "<r.name>:<s.name>",
            <endif>
            }>
            "*"
          ],
          "transactionMode" : "both",
          "searchParam" : [

            <trunc(r.searchParams): {s |
            <\u007B>
              "name" : "<s.name>",
              "type" : "<s.type>"<if(s.targets)>,
              "target" : [
                <trunc(s.targets): {t |
                "<t>",
                }>
                <last(s.targets): {t |
                "<t>"
                }>
              ]
              <endif>
            <\u007D>,
            }>

            <last(r.searchParams): {s |
            <\u007B>
              "name" : "<s.name>",
              "type" : "<s.type>"<if(s.targets)>,
              "target" : [
                <trunc(s.targets): {t |
                "<t>",
                }>
                <last(s.targets): {t |
                "<t>"
                }>
              ]
              <endif>
            <\u007D>
            }>
          ]
        <\u007D>
        }>
      ]
    }
  ]
}