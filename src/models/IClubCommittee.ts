import { ClubsInterface } from "./IClub";
export interface ClubCommitteesInterface {

    ID: number,

    Name: string;

    ID_student: string;

    ClubID: number;
    Club:ClubsInterface;
}