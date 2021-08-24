SELECT
    A.ID,
    A.UserName,
    IF(
        A.Parent > 0,
        (
            SELECT
                B.UserName
            FROM USER B
            WHERE
                B.ID = A.Parent
        ),
        "NULL"
    ) AS ParentUserName
FROM
    USER A;