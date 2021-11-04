import { ActivitiesInterface } from "./IActivity";
import { StudentsInterface } from "./IStudent";
import { ClubCommitteesInterface } from "./IClubCommittee";

export interface JoinActivityHistoryInterface {

    ID: number,

    HourCount: number;

    Point:  number,

    Timestamp: Date;

    ActivityID: number;
    Activity: ActivitiesInterface;

    StudentID: number;
    Student: StudentsInterface;

    EditorID: number;
    Editor: ClubCommitteesInterface;
}