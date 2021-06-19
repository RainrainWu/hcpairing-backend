# hcpairing-backend

The backend server of a hackerthon project for [COVID-19 Healthcare App Challenge](https://healthcareappchallenge.devpost.com/), aims to connect the underserved communities with healthcare provider.

## Getting Started
Please contact the maintainer to obtain the dns name if you wanna give this service a try.

### Search for symptom and feature tags by prefix
```bash=
$ curl -X GET \
  "https://<DNS_NAME>/v1/tags?prefix=s"

{
  "tags":[
    "Sore Muscles",
    "Stomachache"
  ]
}
```

### Convert tags to healthcare scientific name
```bash=
$ curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"zipcode":"52010","tags":["Vomit","Cough"]}' \
  "https://<DNS_NAME>/v1/records"

{
  "specialties":[
    "Pneumology"
  ]
}
```

### Query conversion records by zipcode
```bash=
$ curl -X GET \
  "https://<DNS_NAME>/v1/records?zipcode=52010"

{
  "results":[
    {
      "zipcode":"52010",
      "tags":[
        "Pregnancy",
        "Vomit"
      ]
    }
  ]
}
```

### Search for google map place rating by location name
```bash=
$ curl -X GET \
  "https://<DNS_NAME>/v1/places?name=cheng%20kung"

{
  "name":"National Cheng Kung University",
  "rating":4.6
}
```

## Contributors
- [RainrainWu](https://github.com/RainrainWu)
