const fs = require("fs-extra");
const path = require("path");

UPLOAD_PATH = "public/uploads";
IMAGES_PER_PAGE = 6;

getAllFolders = (req, res) => {
  const uploadPath = path.join(__dirname, UPLOAD_PATH);
  if (!fs.existsSync(uploadPath)) {
    fs.mkdirSync(uploadPath, { recursive: true });
  }
  try {
    const stats = fs.statSync(uploadPath);
    if (!stats.isDirectory()) {
      res.status(500).send("Upload path is not a directory");
    } else {
      const normalizedPath = path.resolve(uploadPath);
      fs.readdir(normalizedPath, (err, folders) => {
        if (err) {
          console.error(err);
          res.status(500).send("Server Error");
        } else {
          res.render("index", { folders: folders });
        }
      });
    }
  } catch (err) {
    res.status(500).send("Error validating upload path:", err);
  }
};

getAllImagesInFolder = (req, res) => {
  const folderName = req.params.folderName;
  var imagesPerPage = req.query.page_size || IMAGES_PER_PAGE;
  imagesPerPage =
    imagesPerPage <= IMAGES_PER_PAGE ? imagesPerPage : IMAGES_PER_PAGE;
  const page_no = req.query.page_no ? parseInt(req.query.page_no) : 1;
  const uploadPath = path.join(__dirname, UPLOAD_PATH, folderName);
  fs.readdir(uploadPath, (err, files) => {
    if (err) {
      console.error(err);
      res.status(500).send("Server Error");
    } else {
      const totalImages = files.length;
      const images = files.slice(
        (page_no - 1) * imagesPerPage,
        page_no * imagesPerPage,
      );
      res.render("show", {
        images: images,
        folderName: folderName,
        currentPage: page_no,
        totalPages: Math.ceil(totalImages / imagesPerPage),
      });
    }
  });
};

module.exports = { getAllFolders, getAllImagesInFolder };
