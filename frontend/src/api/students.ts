import type {
  CreateStudentRequest,
  ListStudentsFilter,
  Student,
  UpdateStudentRequest,
} from "../types/student";
import apiClient from "./axiosClient";

const createStudent = async (req: CreateStudentRequest): Promise<boolean> => {
  const resp = await apiClient.post("/students", req);
  return resp.status === 201;
};

const getStudentByID = async (id: number): Promise<Student> => {
  const resp = await apiClient.get<Student>(`/students/${id}`);
  return resp.data;
};

const listStudents = async (
  filter?: ListStudentsFilter,
): Promise<Student[]> => {
  const params = new URLSearchParams();
  if (filter) {
    if (filter.name) params.append("name", filter.name);
    if (filter.email) params.append("email", filter.email);
    if (typeof filter.group_id !== "undefined")
      params.append("group_id", String(filter.group_id));
    if (typeof filter.limit !== "undefined")
      params.append("limit", String(filter.limit));
    if (typeof filter.offset !== "undefined")
      params.append("offset", String(filter.offset));
  }

  const url = params.toString()
    ? `/students?${params.toString()}`
    : "/students";
  const resp = await apiClient.get<Student[]>(url);
  return resp.data;
};

const getStudentsByGroupID = async (groupId: number): Promise<Student[]> => {
  const resp = await apiClient.get<Student[]>(`/students/group/${groupId}`);
  return resp.data;
};

const updateStudent = async (
  id: number,
  req: UpdateStudentRequest,
): Promise<boolean> => {
  const resp = await apiClient.put(`/students/${id}`, req);
  return resp.status === 200;
};

const deleteStudent = async (id: number): Promise<boolean> => {
  const resp = await apiClient.delete(`/students/${id}`);
  return resp.status === 200 || resp.status === 204;
};

export const studentsApi = {
  createStudent,
  getStudentByID,
  listStudents,
  getStudentsByGroupID,
  updateStudent,
  deleteStudent,
};

export default studentsApi;
