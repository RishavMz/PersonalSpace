const express = require("express");
const multer = require("multer");
const bodyParser = require("body-parser");
const fs = require("fs-extra");
const path = require("path");
const { getAllFolders, getAllImagesInFolder } = require("./controller");

const app = express();
const port = 3000;

app.set("view engine", "ejs");

app.use(bodyParser.urlencoded({ extended: true }));

app.use(express.static("public"));

const storage = multer.diskStorage({
  destination: (req, file, cb) => {
    const folderName = req.body.folderName || "default";
    const uploadPath = path.join(__dirname, "public", "uploads", folderName);
    fs.ensureDirSync(uploadPath);
    cb(null, uploadPath);
  },
  filename: (req, file, cb) => {
    cb(null, Date.now() + path.extname(file.originalname));
  },
});

const upload = multer({ storage: storage });

app.get("/", (req, res) => {
  getAllFolders(req, res);
});

app.post("/upload", upload.array("images", 100), (req, res) => {
  res.redirect("/");
});

app.get("/images/:folderName", (req, res) => {
  getAllImagesInFolder(req, res);
});

app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});
