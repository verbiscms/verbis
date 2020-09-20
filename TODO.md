# To Do

## Fields

Model will::
- Get the fields JSON file from /storage/fields (in paths.Storage)
- You need an endpoint to call to get the admin layout from the fields model which will return below
- You will need a parser to determine whether or not to display the page, via the location block.

```
{
   "status":"ok",
   "status_code":200,
   "data":{
      "friendly_name":"Single Page",
      "options":{
         "extends":"main",
         "canSelect":false
      },
      "fields":[
        {
            "key": "field_5d8c8a2804e95",
            "label": "Title",
            "name": "title",
            "type": "text",
            "instructions": "Add a title for the the resource page.",
            "required": 0,
            "conditional_logic": 0,
            "default_value": "",
            "placeholder": "",
            "maxlength": "",
            "rows": 4
        }
      ]
   },
   "request_time":"2020-09-07 18:31:15"
}
```

```
"location": [
    [
        {
            "param": "page_template",
            "operator": "==",
            "value": "post-single"
        }
    ]
]
```

```php
function compare( $value, $rule ) {
    
    // Allow "all" to match any value.
    if( $rule['value'] === 'all' ) {
        $match = true;
        
    // Compare all other values.
    } else {
        $match = ( $value == $rule['value'] );
    }
    
    // Allow for "!=" operator.
    if( $rule['operator'] == '!=' ) {
        $match = !$match;
    }
    
    // Return.
    return $match;
}
```

## Vue

- Change /resource/editor/2 to just /editor/pageID, have go return if it is a resource or not to avoid the long url in vue.