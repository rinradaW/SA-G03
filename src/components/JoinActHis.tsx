import React, { useEffect } from "react";

import { Link as RouterLink } from "react-router-dom";

import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";

import Typography from "@material-ui/core/Typography";

import Button from "@material-ui/core/Button";

import Container from "@material-ui/core/Container";

import Paper from "@material-ui/core/Paper";

import Box from "@material-ui/core/Box";

import Table from "@material-ui/core/Table";

import TableBody from "@material-ui/core/TableBody";

import TableCell from "@material-ui/core/TableCell";

import TableContainer from "@material-ui/core/TableContainer";

import TableHead from "@material-ui/core/TableHead";

import TableRow from "@material-ui/core/TableRow";

import moment from 'moment';

import { ClubsInterface } from "../models/IClub";

import { ActivitiesInterface } from "../models/IActivity";

import { StudentsInterface } from "../models/IStudent";

import { ClubCommitteesInterface } from "../models/IClubCommittee";

import { JoinActivityHistoryInterface } from "../models/IJoinActivityHistory";

import { format } from 'date-fns'

const useStyles = makeStyles((theme: Theme) =>

  createStyles({

    container: { marginTop: theme.spacing(2) },

    table: { minWidth: 650 },

    tableSpace: { marginTop: 20 },

  })

);



function JoinActivityHistories() {

  const classes = useStyles();

  const [joinActivityHistories, setJoinActivityHistories] = React.useState<JoinActivityHistoryInterface[]>([]);

  const apiUrl = "http://localhost:8080";

  const requestOptions = {

    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },

  };

  
  const getJoinActivityHistories = async () => {
    fetch(`${apiUrl}/join_activity_histories`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setJoinActivityHistories(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    
    getJoinActivityHistories();

  }, []);



  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Box display="flex">
          <Box flexGrow={1}>
          
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              รายการประวัติการเข้าร่วมกิจกรรม
              

            </Typography>
            
          </Box>
          
          <Box>
            <Button
              component={RouterLink}
              
              to="/join_activity_history/create"
              variant="contained"
              color="primary"
            >
              สร้างประวัติกิจกรรม
            </Button>
            </Box>
          
          
          
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="5%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="30%">
                  รายการกิจกรรม
                </TableCell>
                <TableCell align="center" width="15%">
                  ชื่อนักศึกษา
                </TableCell>
                <TableCell align="center" width="5%">
                  ชั่วโมงที่เข้าร่วม
                </TableCell>
                <TableCell align="center" width="5%">
                  คะแนน
                </TableCell>
                <TableCell align="center" width="20%">
                  ผู้แก้ไข
                </TableCell>
                <TableCell align="center" width="20%">
                  วันที่และเวลา
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {joinActivityHistories.map((item: JoinActivityHistoryInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center" >{item.ID}</TableCell>
                  <TableCell align="center" >{item.Activity.Name}</TableCell>
                  <TableCell align="center" >{item.Student.Name}</TableCell>
                  <TableCell align="center" >{item.HourCount}</TableCell>
                  <TableCell align="center">{item.Point}</TableCell>
                  <TableCell align="center">{item.Editor.Name}</TableCell>
                  <TableCell align="center">{format((new Date(item.Timestamp)), 'dd MMMM yyyy hh:mm a')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default JoinActivityHistories;