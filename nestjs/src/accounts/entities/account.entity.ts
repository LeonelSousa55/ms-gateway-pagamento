import {
    Column,
    DataType,
    Model,
    PrimaryKey,
    Table
} from "sequelize-typescript";

@Table({
    tableName: 'accounts',
    createdAt: 'created_at',
    updatedAt: 'updated_ap',
})

export class Account extends Model {

    @PrimaryKey
    @Column({ type: DataType.UUID, defaultValue: DataType.UUIDV4 })
    id: string;

    @Column({ allowNull: false })
    name: string;

    @Column({
        allowNull: false,
        defaultValue: () => Math.random().toString(36).slice(2),
    })
    token: string;
}