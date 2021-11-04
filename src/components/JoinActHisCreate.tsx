import { useEffect, useState } from "react";

import { Link as RouterLink } from "react-router-dom";

import { makeStyles, Theme, createStyles } from "@material-ui/core/styles";

import TextField from "@material-ui/core/TextField";

import Button from "@material-ui/core/Button";

import FormControl from "@material-ui/core/FormControl";

import Container from "@material-ui/core/Container";

import Paper from "@material-ui/core/Paper";

import Grid from "@material-ui/core/Grid";

import Box from "@material-ui/core/Box";

import Typography from "@material-ui/core/Typography";

import Divider from "@material-ui/core/Divider";

import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";

import Snackbar from "@material-ui/core/Snackbar";

import { MuiPickersUtilsProvider, KeyboardDateTimePicker, } from "@material-ui/pickers";

import DateFnsUtils from "@date-io/date-fns";

import Select from '@material-ui/core/Select';

import InputLabel from '@material-ui/core/InputLabel';

import { FormHelperText } from "@material-ui/core";

import { ClubsInterface } from "../models/IClub";

import { ActivitiesInterface } from "../models/IActivity";

import { StudentsInterface } from "../models/IStudent";

import { ClubCommitteesInterface } from "../models/IClubCommittee";

import { JoinActivityHistoryInterface } from "../models/IJoinActivityHistory";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};


const useStyles = makeStyles((theme: Theme) =>

  createStyles({

    root: { flexGrow: 1 },

    container: { marginTop: theme.spacing(8) },

    paper: { padding: theme.spacing(2), color: theme.palette.text.secondary },

    formControl: {
      margin: theme.spacing(0),
      minWidth: 260,
    },

    buttonControl: {
      margin: theme.spacing(1),
      minWidth: 80,
    },

    selectEmpty: {
      marginTop: theme.spacing(2),
    },

    margin: {
      margin: theme.spacing(1),
      minWidth: 120,
    },

  })

);

function JoinActivityHistoryCreate() {

  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());

  const [ClubCommittees, setClubCommittees] = useState<ClubCommitteesInterface>();
  const [Clubs, setClubs] = useState<ClubsInterface[]>([]);
  const [Activities, setActivities] = useState<ActivitiesInterface[]>([]);
  const [Students, setStudents] = useState<StudentsInterface[]>([]);
  const [JoinActivityHistory, setJoinActivityHistory] = useState<Partial<JoinActivityHistoryInterface>>(
    {}
  );

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const apiUrl = "http://localhost:8080";

  

  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
      
    },
  };

  
  {/*console.log(loginUser.getItem("uid"));*/}

  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof JoinActivityHistory;
    setJoinActivityHistory({
      ...JoinActivityHistory,
      [name]: event.target.value,
    });
    
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  

  const getClubCommittees = async () => {
    let uid = localStorage.getItem("uid");
    fetch(`${apiUrl}/club_committee/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        JoinActivityHistory.EditorID = res.data.ID
        if (res.data) {
          setClubCommittees(res.data);
          {/*console.log(res.data);*/}
        } else {
          console.log("else");
        }
      });
  };

  const getClubs = async () => {
    fetch(`${apiUrl}/clubs`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setClubs(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getActivities = async () => {
    fetch(`${apiUrl}/activities`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setActivities(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getStudents = async () => {
    fetch(`${apiUrl}/students`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setStudents(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getClubCommittees();
    getClubs();
    getActivities();
    getStudents();
  }, []);

  
  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      ActivityID: convertType(JoinActivityHistory?.ActivityID),
      StudentID: convertType(JoinActivityHistory?.StudentID),
      HourCount: convertType(JoinActivityHistory?.HourCount),
      Point: convertType(JoinActivityHistory?.Point),
      /*edit around here*/
      EditorID: convertType(JoinActivityHistory?.EditorID),
      Timestamp: selectedDate,
    };

    const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/join_activity_histories`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true);
        } else {
          console.log("บันทึกไม่ได้")
          setError(true);
        }
      });
  }

  console.log(JoinActivityHistory);

  return (

    <Container className={classes.container} maxWidth="sm">
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>

        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>

      </Snackbar>

      <Paper className={classes.paper}>

        <Box display="flex">

          <Box flexGrow={1}>

            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกประวัติการเข้าร่วมกิจกรรมชมรม
            </Typography>

          </Box>
        </Box>

        <Divider />
        <Typography component="div" style={{ height: '4vh' }} />
        <Grid container spacing={2} className={classes.root}>

          <Grid item xs={5}>

            <p>กิจกรรมที่เข้าร่วม</p>
          </Grid>
          <Grid item xs={7}>

            <FormControl variant="outlined" fullWidth>
              <Select
                native
                value={JoinActivityHistory.ActivityID}
                onChange={handleChange}
                inputProps={{
                  name: "ActivityID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกกิจกรรมที่เข้าร่วม
                </option>
                {Activities.map((item: ActivitiesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>

            </FormControl>

          </Grid>



          <Grid item xs={5}>
            <p>สมาชิกที่เข้าร่วม</p>
          </Grid>
          <Grid item xs={7}>
            <FormControl fullWidth variant="outlined">

              <Select
                native
                value={JoinActivityHistory.StudentID}
                onChange={handleChange}
                inputProps={{
                  name: "StudentID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกสมาชิกที่เข้าร่วม
                </option>
                {Students.map((item: StudentsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>

            </FormControl>



          </Grid>

          <Grid item xs={5}>
            <p>รายชั่วโมงที่เข้าร่วม</p>
          </Grid>

          <Grid item xs={7}>
            <FormControl fullWidth variant="outlined">

              <TextField
                id="HourCount"
                variant="outlined"

                type="integer"

                size="medium"

                inputProps={{
                  name: 'HourCount',
                }}
                 InputLabelProps={{
                   shrink: true,
                 }}

                onChange={handleChange}

              />

            </FormControl>
          </Grid>

          <Grid item xs={5}>
            <p>คะแนนจิตอาสา</p>
          </Grid>

          <Grid item xs={7}>
            <FormControl fullWidth variant="outlined">

              <TextField
                id="Point"
                variant="outlined"

                type="integer"

                size="medium"

                inputProps={{
                  name: 'Point',
                }}
                 InputLabelProps={{
                   shrink: true,
                 }}

                onChange={handleChange}

              />

            </FormControl>
          </Grid>

          <Grid item xs={5}>
            <p>ผู้แก้ไข</p>
          </Grid>
          <Grid item xs={7}>
            <FormControl variant="outlined" className={classes.formControl} fullWidth disabled>

              <InputLabel htmlFor="outlined-edited-native-simple"></InputLabel>
              <Select
                native
                value={JoinActivityHistory.EditorID}
                
                inputProps={{
                  name: "EditorID",
                }}
              >
                {/*edit around here*/}
                
                <option value={ClubCommittees?.ID} key={ClubCommittees?.ID}>
                  {ClubCommittees?.Name}
                </option>

                {/* {ClubCommittees.map((item: ClubCommitteesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))} */}
                
                
              </Select>

              <FormHelperText>* disabled </FormHelperText>
            </FormControl>

          </Grid>

          <Grid item xs={5}>
            <p>ตราประทับเวลา</p>
          </Grid>
          <Grid item xs={7}>
            <FormControl variant="outlined" fullWidth >

              <MuiPickersUtilsProvider utils={DateFnsUtils}>

              <KeyboardDateTimePicker disabled
                  name="Timestamp"
                  value={selectedDate}
                  onChange={handleDateChange}
                  label="กรุณาเลือกวันที่และเวลา"
                  minDate={new Date("2018-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                />

              </MuiPickersUtilsProvider>

            </FormControl>

          </Grid>

          <Grid item xs={12}>

            <Button className={classes.buttonControl}
              component={RouterLink} to="/join_activity_histories"
              variant="contained" size="large"
              style={{ float: "left" }}>
              Back
            </Button>

            <Button className={classes.buttonControl}
              style={{ float: "right" }}
              variant="contained"
              color="primary"
              size="large"
              onClick={submit}
            >

              บันทึกประวัติการเข้าร่วมกิจกรรม

            </Button>

          </Grid>

        </Grid>

      </Paper>

      <Typography component="div" style={{ height: '15vh' }} />

    </Container>
  );

}
export default JoinActivityHistoryCreate;