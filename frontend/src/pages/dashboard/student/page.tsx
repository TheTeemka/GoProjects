import { useEffect, useState } from "react";
import { StudentTable } from "@/components/students/StudentTable";
import { Button } from "@/components/ui/button";
import { Plus, Search } from "lucide-react";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { useStudents } from "@/hooks/useStudents";
import AddStudentModal from "@/components/students/AddStudentModal";
import UpdateStudentModal from "@/components/students/UpdateStudentModal";
import DeleteStudentModal from "@/components/students/DeleteStudentModal";
import type { Student } from "@/types/student";

export default function StudentsPage() {
  const [searchTerm, setSearchTerm] = useState<string>("");
  const { students, applyFilter, createStudent, updateStudent, deleteStudent } =
    useStudents({
      group_id: 1,
    });
  const [editingStudent, setEditingStudent] = useState<Student | null>(null);
  const [deletingId, setDeletingId] = useState<number | null>(null);

  useEffect(() => {
    const id = setTimeout(() => {
      applyFilter({
        name: searchTerm || undefined,
        email: searchTerm || undefined,
      });
    }, 250);
    return () => clearTimeout(id);
  }, [searchTerm, applyFilter]);

  return (
    <div className="space-y-6 px-10 py-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Students</h1>
          <p className="text-gray-500 mt-2">Manage your students</p>
        </div>
      </div>

      <div className="py-6 grid grid-cols-3 gap-10">
        <Card>
          <CardHeader>
            <CardTitle>Group ID</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-2xl font-semibold">1</p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Number of Students</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-2xl font-semibold">{students.length}</p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Attendance Percentage</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-2xl font-semibold">MOCK</p>
          </CardContent>
        </Card>
      </div>

      <div className="flex items-center justify-between gap-4">
        <div className="flex-1 grow flex items-center bg-white dark:bg-gray-800 px-4 py-2 rounded-lg border">
          <div className="pr-4">
            <Search className="h-4 w-4 text-gray-400" />
          </div>
          <input
            type="text"
            placeholder="Search by name or email..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="flex-1 outline-none bg-transparent"
          />
        </div>
        <div className="flex">
          <AddStudentModal
            defaultGroupId={1}
            trigger={
              <Button className="ml-auto">
                <Plus className="mr-2 h-4 w-4" />
                Add Student
              </Button>
            }
            onConfirm={createStudent}
            onSuccess={() => {
              /* hook already refreshes */
            }}
          />
        </div>
      </div>
      <StudentTable
        students={students}
        onEdit={(s) => setEditingStudent(s)}
        onDelete={(id) => setDeletingId(id)}
      />

      <UpdateStudentModal
        student={editingStudent}
        onConfirm={updateStudent}
        onSuccess={() => setEditingStudent(null)}
        onClose={() => setEditingStudent(null)}
      />

      <DeleteStudentModal
        id={deletingId}
        onConfirm={deleteStudent}
        onSuccess={() => setDeletingId(null)}
        onClose={() => setDeletingId(null)}
      />
    </div>
  );
}
