export type Student = {
  id: number;
  name: string;
  birthday: string;
  email: string;
  group_id: number;
};

export type CreateStudentRequest = Omit<Student, "id">;

export type UpdateStudentRequest = Partial<CreateStudentRequest>;

export type ListStudentsFilter = {
  name?: string;
  email?: string;
  group_id?: number;
  limit?: number;
  offset?: number;
};
