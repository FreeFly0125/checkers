import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgPlayMove } from "./types/mycheckers/mycheckers/tx";
import { MsgCreateGame } from "./types/mycheckers/mycheckers/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/mycheckers.mycheckers.MsgPlayMove", MsgPlayMove],
    ["/mycheckers.mycheckers.MsgCreateGame", MsgCreateGame],
    
];

export { msgTypes }