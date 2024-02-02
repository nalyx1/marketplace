import { Type } from 'class-transformer';
import {
  ArrayNotEmpty,
  IsInt,
  IsNotEmpty,
  IsPositive,
  IsString,
  MaxLength,
  ValidateNested,
} from 'class-validator';

export class CreateOrderDto {
  @ValidateNested({ each: true })
  @ArrayNotEmpty()
  @Type(() => OrderItemDto)
  items: OrderItemDto[];

  @IsNotEmpty()
  @IsString()
  @MaxLength(255)
  card_hash: string;
}

export class OrderItemDto {
  @IsInt()
  @IsNotEmpty()
  @IsPositive()
  quantity: number;

  @MaxLength(36)
  @IsString()
  @IsNotEmpty()
  product_id: string;
}
