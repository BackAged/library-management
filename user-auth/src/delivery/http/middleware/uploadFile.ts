import multer from "multer";
import path from "path";

const storage = multer.diskStorage({
    destination: (req, file, cb) => {
      cb(null, path.join(__dirname, '../../../../storage/'));
    },

    filename: (req, file, cb) => {
      cb(null, `${file.fieldname}_${+new Date()}.jpg`);
    }
});

export const upload = multer({
    storage,
    fileFilter: (req, file, cb) => {
      if (
        !file.mimetype.includes("jpeg") &&
        !file.mimetype.includes("jpg") &&
        !file.mimetype.includes("png") &&
        !file.mimetype.includes("gif")
      ) {
        return cb(new Error("Only images are allowed"));
      }
      cb(null, true);
    }
});