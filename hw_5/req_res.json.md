# Attendance Route

## POST {{baseURL}}/attendance/subject
### Req
```json
{
  "student_id": 1,
  "subject_id": 1,
  "visit_date": "2026-01-13",
  "visited": false
}
```

### Res
```json
{
    "message": "Attendance record created successfully"
}
```

## GET {{baseURL}}/attendance/attendanceByStudentId/:student_id
### Res
```json
[
    {
        "id": 1,
        "student_id": 1,
        "subject_id": 101,
        "visit_date": "2023-10-01",
        "visited": true
    },
    {
        "id": 2,
        "student_id": 1,
        "subject_id": 102,
        "visit_date": "2023-10-02",
        "visited": false
    },
    {
        "id": 3,
        "student_id": 1,
        "subject_id": 1,
        "visit_date": "2026-01-13",
        "visited": false
    }
]
```

## GET {{baseURL}}/attendance/attendanceBySubjectId/:subject_id

### Res
```json
[
    {
        "id": 1,
        "student_id": 1,
        "subject_id": 101,
        "visit_date": "2023-10-01",
        "visited": true
    },
    {
        "id": 4,
        "student_id": 1,
        "subject_id": 101,
        "visit_date": "2026-01-13",
        "visited": false
    }
]
```

# Student Route
## GET {{baseURL}}/schedule/student/:student_id
### Res
```json
[
    {
        "ID": 1,
        "Subject": "Math",
        "DayOfWeek": "Monday",
        "Time": "9:00",
        "GroupID": 1
    },
    {
        "ID": 2,
        "Subject": "Physics",
        "DayOfWeek": "Tuesday",
        "Time": "10:00",
        "GroupID": 1
    },
    {
        "ID": 3,
        "Subject": "English",
        "DayOfWeek": "Wednesday",
        "Time": "11:00",
        "GroupID": 1
    },
    {
        "ID": 4,
        "Subject": "History",
        "DayOfWeek": "Thursday",
        "Time": "12:00",
        "GroupID": 1
    },
    {
        "ID": 5,
        "Subject": "Chemistry",
        "DayOfWeek": "Friday",
        "Time": "13:00",
        "GroupID": 1
    }
]
```

## GET {{baseURL}}/schedule/schedule/group/:id
### Res 
```json
[
    {
        "ID": 1,
        "Subject": "Math",
        "DayOfWeek": "Monday",
        "Time": "9:00",
        "GroupID": 1
    },
    {
        "ID": 2,
        "Subject": "Physics",
        "DayOfWeek": "Tuesday",
        "Time": "10:00",
        "GroupID": 1
    },
    {
        "ID": 3,
        "Subject": "English",
        "DayOfWeek": "Wednesday",
        "Time": "11:00",
        "GroupID": 1
    },
    {
        "ID": 4,
        "Subject": "History",
        "DayOfWeek": "Thursday",
        "Time": "12:00",
        "GroupID": 1
    },
    {
        "ID": 5,
        "Subject": "Chemistry",
        "DayOfWeek": "Friday",
        "Time": "13:00",
        "GroupID": 1
    }
]
```

## GET {{baseURL}}/schedule/all_class_schedule
### Res
```json
[
    {
        "ID": 1,
        "Subject": "Math",
        "DayOfWeek": "Monday",
        "Time": "9:00",
        "GroupID": 1
    },
    {
        "ID": 2,
        "Subject": "Physics",
        "DayOfWeek": "Tuesday",
        "Time": "10:00",
        "GroupID": 1
    },
    {
        "ID": 3,
        "Subject": "English",
        "DayOfWeek": "Wednesday",
        "Time": "11:00",
        "GroupID": 1
    },
    {
        "ID": 4,
        "Subject": "History",
        "DayOfWeek": "Thursday",
        "Time": "12:00",
        "GroupID": 1
    },
    {
        "ID": 5,
        "Subject": "Chemistry",
        "DayOfWeek": "Friday",
        "Time": "13:00",
        "GroupID": 1
    },
    {
        "ID": 6,
        "Subject": "Biology",
        "DayOfWeek": "Monday",
        "Time": "9:00",
        "GroupID": 2
    },
    {
        "ID": 7,
        "Subject": "Algebra",
        "DayOfWeek": "Tuesday",
        "Time": "10:00",
        "GroupID": 2
    },
    {
        "ID": 8,
        "Subject": "Literature",
        "DayOfWeek": "Wednesday",
        "Time": "11:00",
        "GroupID": 2
    },
    {
        "ID": 9,
        "Subject": "Geography",
        "DayOfWeek": "Thursday",
        "Time": "12:00",
        "GroupID": 2
    },
    {
        "ID": 10,
        "Subject": "Computer Science",
        "DayOfWeek": "Friday",
        "Time": "13:00",
        "GroupID": 2
    },
    {
        "ID": 11,
        "Subject": "Calculus",
        "DayOfWeek": "Monday",
        "Time": "9:00",
        "GroupID": 3
    },
    {
        "ID": 12,
        "Subject": "Physics",
        "DayOfWeek": "Tuesday",
        "Time": "10:00",
        "GroupID": 3
    },
    {
        "ID": 13,
        "Subject": "English",
        "DayOfWeek": "Wednesday",
        "Time": "11:00",
        "GroupID": 3
    },
    {
        "ID": 14,
        "Subject": "History",
        "DayOfWeek": "Thursday",
        "Time": "12:00",
        "GroupID": 3
    },
    {
        "ID": 15,
        "Subject": "Art",
        "DayOfWeek": "Friday",
        "Time": "13:00",
        "GroupID": 3
    },
    {
        "ID": 16,
        "Subject": "Math",
        "DayOfWeek": "Monday",
        "Time": "9:00",
        "GroupID": 4
    },
    {
        "ID": 17,
        "Subject": "Science",
        "DayOfWeek": "Tuesday",
        "Time": "10:00",
        "GroupID": 4
    },
    {
        "ID": 18,
        "Subject": "Literature",
        "DayOfWeek": "Wednesday",
        "Time": "11:00",
        "GroupID": 4
    },
    {
        "ID": 19,
        "Subject": "Economics",
        "DayOfWeek": "Thursday",
        "Time": "12:00",
        "GroupID": 4
    }
]
```