# GSON

<https://github.com/albertattard/gson-typeadapter-example>

## TypeAdapter

```kt
import com.google.gson.JsonSyntaxException
import com.google.gson.TypeAdapter
import com.google.gson.stream.JsonReader
import com.google.gson.stream.JsonToken
import com.google.gson.stream.JsonWriter


class NumberTypeAdapter : TypeAdapter<Number?>() {
    override fun write(jsonWriter: JsonWriter, number: Number?) {
        if (number == null) jsonWriter.nullValue()
        else jsonWriter.value(number)
    }

    override fun read(jsonReader: JsonReader): Number? {
        if (jsonReader.peek() == JsonToken.NULL) {
            jsonReader.nextNull()
            return null
        }

        return try {
            val value = jsonReader.nextString()
            if (value == "") 0 else value.toInt()
        } catch (e: NumberFormatException) {
            throw JsonSyntaxException(e)
        }
    }
}
```

사용

```kt
val gson = GsonBuilder()
        .setFieldNamingStrategy(FieldNamingPolicy.UPPER_CAMEL_CASE)
        .registerTypeAdapter(Long::class.java, NumberTypeAdapter())
        .create()
```
