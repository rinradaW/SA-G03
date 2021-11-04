import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function Home() {
  const classes = useStyles();

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบบันทึกประวัติการเข้าร่วมกิจกรรม</h1>
        <h4>Requirements</h4>
        <p>
        ระบบบันทึกประวัติการเข้าร่วมกิจกรรม เป็นระบบที่กรรมการบริหารชมรมที่ทำการจัดกิจกรรม
        เป็นผู้รับรองการเข้าร่วมกิจกรรม และเพิ่มประวัติการเข้าร่วมกิจกรรมชมรมเข้าไปในระบบ
        ระบบจะทำการบันทึกประวัติการเข้าร่วมกิจกรรม เพื่อเป็นประวัติย้อนหลังให้กรรมการบริหารชมรมสามารถสืบค้น
        และเพื่อเป็นหลักฐานยืนยันการเข้าร่วมกิจกรรมของสมาชิก ให้กรรมการบริหารชมรม
        ยื่นเรื่องเพิ่มคะแนนจิตอาสาให้กับบุคคลที่เข้าร่วมกิจกรรมให้หน่วยงานที่เกี่ยวข้องต่อไป
        </p>
      </Container>
    </div>
  );
}
export default Home;