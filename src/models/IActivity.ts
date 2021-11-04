import { ClubsInterface } from "./IClub";
export interface ActivitiesInterface {

    ID: number,

    Name: string;

    Time: Date;

    Amount: number;

    ClubID: number;
    Club: ClubsInterface;
}