const asyncHandler = require("express-async-handler");
const { getAllStudents, addNewStudent, getStudentDetail, setStudentStatus, updateStudent } = require("./students-service");

const handleGetAllStudents = asyncHandler(async (req, res) => {
    //write your code

    const students = await getAllStudents(req.query);

    res.status(200).json({
        success: true,
        count: students.length,
        data: students
    });

});

const handleAddStudent = asyncHandler(async (req, res) => {
    //write your code

});

const handleUpdateStudent = asyncHandler(async (req, res) => {
    //write your code

});

const handleGetStudentDetail = asyncHandler(async (req, res) => {
    //write your code
    const { id } = req.params;

    const student = await getStudentDetail(id);
    res.status(200).json({
        success: true,
        data: student
    });
});

const handleStudentStatus = asyncHandler(async (req, res) => {
    //write your code

});

module.exports = {
    handleGetAllStudents,
    handleGetStudentDetail,
    handleAddStudent,
    handleStudentStatus,
    handleUpdateStudent,
};
