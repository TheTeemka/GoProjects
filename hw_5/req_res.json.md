1. Student Route
GET {{baseURL}}/student/:id

### Res
```json
{
  "id": 1,
  "name": "Ivan Ivanov",
  "birth_date": "2004-05-15T00:00:00Z",
  "gender": "male",
  "group_id": 10,
  "group_name": "SE-2201"
}
```

## GET {{baseURL}}/schedule/all_class_schedule
### Res
```json
[
    {
        "id": 1,
        "group_id": 10,
        "subject": "Mathematics",
        "start_time": "09:00",
        "end_time": "10:30"
    },
    {
        "id": 2,
        "group_id": 11,
        "subject": "Physics",
        "start_time": "11:00",
        "end_time": "12:30"
    }
]
```

## GET {{baseURL}}/schedule/group/:id
### Res
```json
[
    {
        "id": 5,
        "group_id": 1,
        "subject": "Golang Backend",
        "start_time": "14:00",
        "end_time": "15:30"
    }
]
```

# Attendance Route

## POST {{baseURL}}/attendance/subject
### Req
```json
{
  "subject_id": 101,
  "visit_day": "25.10.2023", 
  "visited": true,
  "student_id": 55
}
```

### Res
```json
{
  "id": 12
}
```

## GET {{baseURL}}/attendance/attendanceBySubjectId/:subject_id
### Res
```json
[
    {
        "id": 1,
        "subject_id": 101,
        "visit_day": "2023-10-25T00:00:00Z",
        "visited": true,
        "student_id": 55
    },
    {
        "id": 2,
        "subject_id": 101,
        "visit_day": "2023-10-25T00:00:00Z",
        "visited": false,
        "student_id": 56
    }
]
```

## GET {{baseURL}}/attendance/attendanceByStudentId/:student_id
### Res
```json
[
    {
        "id": 10,
        "subject_id": 101,
        "visit_day": "2023-10-25T00:00:00Z",
        "visited": true,
        "student_id": 55
    },
    {
        "id": 25,
        "subject_id": 102,
        "visit_day": "2023-10-26T00:00:00Z",
        "visited": true,
        "student_id": 55
    }
]
```