**index**

+ ngoài cùng bên trái

+ trường nào có tỷ lệ lặp dữ liệu ít nhất thì đánh index nằm ngoài cùng rồi dựa vào tỷ lệ mà sắp xếp đánh index

+ Truy vấn lấy trong khoảng là mất index

**Json trong mysql**

**DEADLOCK trong mysql(Liên quan đến transactions)**

**Khóa mutex(Lập trình lượng đồng thời cao bảo toàn dữ liệu)**

**debezium(nên dùng dễ cấu hình và hỗ trợ nhiều loại database)**

**canal(chỉ hỗ trợ mysql)**

**Dùng EXPLAIN để đánh giá chất lượng câu truy vấn (TYPE)**

system -> const -> eq_ref -> ref -> range -> index -> ALL (sắp xếp chất lượng câu truy vấn dữ liệu từ cao đến thấp)

Khi nào sử dụng mongoDB



**tối ưu lệnh join trong query**

(Tương ứng với dùng vòng lặp qua các kiều kiện)

+ thứ nhất phải chọn bảng control có ít dữ liệu nhất

+ thứ 2 phải where tối ưu nhất dựa vào Type (system -> const -> eq_ref -> ref -> range -> index -> ALL)

**Tối ưu bằng truy vấn phụ(4 cú pháp)**

**The subquery as scalar Operand**

ex: SELECT (SELECT i1 FROM t1 LIMIT 1); => truy vấn này trả về 1 giá trị duy nhất

**Row SubQueries**

ex: SELECT * FROM t1 WHERE (i1,j1) = (SELECT i2,j2 FROM t2 LIMIT 1) => **trả về nhiều trường nhưng chỉ có 1 row**

**Column SubQueries**

ex: SELECT * FROM t1 WHERE i1 IN (SELECT i2 FROM t2); => **truy vấn con trả về 1 trường dữ liệu nhưng chứa nhiều bảng ghi** 

**Table Subqueries**

ex: SELECT * FROM t1 WHERE (i1,j1) IN (SELECT i2,j2 FROM t2) => **trả về nhiều trường và nhiều row**

**Khấu trừ hàng tồn kho**

Tạo storeprocedure

database

```
| Id | Description | stock_initial | is_stock_prepared | price_original | price_flash | sale_start_time       | sale_end_time         | status | activite_id | update_at           | created_at          | stock_version |
|----|------------|--------------|------------------|---------------|------------|---------------------|---------------------|--------|------------|---------------------|---------------------|--------------|
| 3  | vé sự kiện  | 2000         | 2000             | 0             | 100000     | 2025-01-01 00:00:00 | 2025-01-01 00:00:00 | 1      | 2          | 2025-02-15 03:37:26 | 2025-02-05 08:17:01 | 0            |

```

```
DELIMITER
CREATE PROCEDURE descrese_stock_with_cas(
    INT p_item_id BIGINT,
    INT p_quantity INT,
    INT p_old_stock_available INT
)
BEGIN 
    UPDATE ticket_item
    SET stock_available = p_old_stock_available - p_quantity,
        update_at = NOW()
    WHERE id = p_item_id
     AND stock_available = p_old_stock_available;
END
DELIMITER;
```

**jit đại diện cho 1 token riêng biệt**

Nhằm mục đích phân biệt logout cho 1 tài trên từng thiết bị

**Solution architecture tính nhất quán dữ liệu trong localcache**

```
Tạo 1 trường version để so sánh nếu 2 version không giống nhau dữ liệu đã thay đổi còn nếu version giống nhau thì dữ liệu được nhất quán.
```

Lập trình LUA trong redis tương tự như transaction trong mysql.Các lệnh LUA trong redis

```
MULTI: mở 1 giao dịch đánh dấu mở 1 transaction trong redis lệnh này không thực thi ngay lập tức mà đưa vào queues cho đến khi chúng ta EXEC và tất cả câu lệnh thực thi.
DISCARD: hủy giao dịch 
EXEC: thực thi tất cả câu lệnh trong hàm đợi 
WATCH  
UNWATCH 
```

Search trong mysql

```
CREATE FULLTEXT INDEX idx_role_name ON role(role_name);
SELECT * FROM products WHERE MATCH(productName) AGAINST ('suzuki')
SELECT * FROM products WHERE MATCH(productName) AGAINST ('suzuki corsair') // tìm kiếm gần đúng
SELECT * FROM products WHERE MATCH(productName) AGAINST ('+suzuki +corsair' IN BOOLEAN MODE) // đúng chính xác
```
