import studentsApi from "@/api/students";
import type {
  Student,
  CreateStudentRequest,
  ListStudentsFilter,
} from "@/types/student";
import { useCallback, useEffect, useRef, useState } from "react";

export function useStudents(initialFilter?: ListStudentsFilter) {
  const [filter, setFilter] = useState<ListStudentsFilter | undefined>(
    initialFilter,
  );
  const filterRef = useRef<ListStudentsFilter | undefined>(filter);
  useEffect(() => {
    filterRef.current = filter;
  }, [filter]);
  const [students, setStudents] = useState<Student[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const fetchStudents = useCallback(
    async (overrideFilter?: ListStudentsFilter) => {
      setLoading(true);
      setError(null);
      try {
        const f = overrideFilter ?? filterRef.current;
        const data = await studentsApi.listStudents(f);
        setStudents(data);
      } catch (err) {
        setError(
          err instanceof Error ? err.message : "Failed to fetch students",
        );
      } finally {
        setLoading(false);
      }
    },
    [],
  );

  useEffect(() => {
    fetchStudents();
  }, [filter]);

  const applyFilter = useCallback((updater: Partial<ListStudentsFilter>) => {
    const nextFilter = { ...(filter ?? {}), ...updater };
    setFilter(nextFilter);
  }, []);

  const createStudent = async (
    student: CreateStudentRequest,
  ): Promise<boolean> => {
    setLoading(true);
    setError(null);
    try {
      const success = await studentsApi.createStudent(student);
      if (success) {
        await fetchStudents();
        return true;
      } else {
        setError("Failed to create student");
        return false;
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to create student");
      return false;
    } finally {
      setLoading(false);
    }
  };

  const updateStudent = async (
    id: number,
    student: Partial<CreateStudentRequest>,
  ): Promise<boolean> => {
    setLoading(true);
    setError(null);
    try {
      const success = await studentsApi.updateStudent(id, student);
      if (success) {
        await fetchStudents();
        return true;
      } else {
        setError("Failed to update student");
        return false;
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to update student");
      return false;
    } finally {
      setLoading(false);
    }
  };

  const deleteStudent = async (id: number): Promise<boolean> => {
    setLoading(true);
    setError(null);
    try {
      const success = await studentsApi.deleteStudent(id);
      if (success) {
        await fetchStudents();
        return true;
      } else {
        setError("Failed to delete student");
        return false;
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to delete student");
      return false;
    } finally {
      setLoading(false);
    }
  };

  return {
    students,
    loading,
    error,
    filter,
    setFilter,
    applyFilter,
    fetchStudents,
    createStudent,
    updateStudent,
    deleteStudent,
  };
}
