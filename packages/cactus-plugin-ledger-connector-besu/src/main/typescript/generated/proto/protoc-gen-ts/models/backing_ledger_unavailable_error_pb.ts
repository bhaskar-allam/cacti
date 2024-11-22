/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.19.1
 * source: models/backing_ledger_unavailable_error_pb.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../google/protobuf/any";
import * as pb_1 from "google-protobuf";
export namespace org.hyperledger.cacti.plugin.ledger.connector.besu {
    export class BackingLedgerUnavailableErrorPB extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            message?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("message" in data && data.message != undefined) {
                    this.message = data.message;
                }
            }
        }
        get message() {
            return pb_1.Message.getFieldWithDefault(this, 418054152, "") as string;
        }
        set message(value: string) {
            pb_1.Message.setField(this, 418054152, value);
        }
        static fromObject(data: {
            message?: string;
        }): BackingLedgerUnavailableErrorPB {
            const message = new BackingLedgerUnavailableErrorPB({});
            if (data.message != null) {
                message.message = data.message;
            }
            return message;
        }
        toObject() {
            const data: {
                message?: string;
            } = {};
            if (this.message != null) {
                data.message = this.message;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.message.length)
                writer.writeString(418054152, this.message);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): BackingLedgerUnavailableErrorPB {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new BackingLedgerUnavailableErrorPB();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 418054152:
                        message.message = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): BackingLedgerUnavailableErrorPB {
            return BackingLedgerUnavailableErrorPB.deserialize(bytes);
        }
    }
}
