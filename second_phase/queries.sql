--.1 سه استیکر پر استفاده کاربر با شناسه mohammad را چاپ کنید.

SELECT *
FROM stickers
WHERE id IN (
    SELECT sticker_id
    FROM account_stickers
    WHERE account_id IN (
        SELECT id
        FROM accounts
        WHERE username = 'Mohammad'
    )
)
ORDER BY id DESC
LIMIT 3;
--.2 متام خماطبان مشرتک بین کاربر با شناسه ali و mohammad را چاپ کنید.
SELECT *
FROM accounts
WHERE id IN (
    SELECT contact_id
    FROM contacts
    WHERE account_id IN (
        SELECT id
        FROM accounts
        WHERE username = 'Ali'
    )
      AND contact_id IN (
        SELECT contact_id
        FROM contacts
        WHERE account_id IN (
            SELECT id
            FROM accounts
            WHERE username = 'Mohammad'
        )
    )
);

-- تعداد پیام ها به همراه نام گروه های با تعداد کاربر بیش از ۱۰ هزار نفر را به ترتیب تعداد چاپ کنید
SELECT
    g.name AS group_name,
    (
        SELECT COUNT(*)
        FROM messages m
        WHERE m.chat_id = g.id
    )
FROM
    groups g
        JOIN
    chats c ON g.id = c.id
WHERE
    g.members_count > 10000
ORDER BY
    messages_count DESC;


------------------------------------
--.4 نام گروه با بیشرتین تعداد پیام در ۱۰ روز اخیر را چاپ کنید.
SELECT g.name
FROM groups g
         JOIN chats c ON g.id = c.id
         JOIN messages m ON c.id = m.chat_id
WHERE m.created_at >= NOW() - INTERVAL '10 days'
GROUP BY g.name
ORDER BY COUNT(m.id) DESC
LIMIT 1;


SELECT name
FROM groups
WHERE id IN (
    SELECT id
    FROM chats
    WHERE chat_type = 'group'
    ORDER BY messages_count DESC
    LIMIT 1
);

---------------------------------------
--.5 جمموع حجم فایل های آپلود شده توسط کاربر با بیشرتین عضویت در گروه ها را چاپ کنید.

select sum(file_size) from storage
group by uploader_id
having uploader_id in (
    select account_id from group_members
    group by account_id
    order by count(group_id)
            desc limit 1
);
----------------------
--.6 نام کاربرانی که حداکثر در دو گروه عضو هستند را چاپ کنید.
select username from accounts
where id not in (
    select account_id from group_members
                      group by account_id having count(group_id)>=3
    );


---------------------------------
--.7 مشاره متاس افرادی که با ۰۹۱۲ شروع و به ۸ ختم میشود را چاپ کنید.
SELECT mobile
FROM accounts
WHERE mobile LIKE '0912%8';
------------------------------------
--.8 کاربرانی که دقیقا دارای ۲ کانال هستند را لیست کنید.
SELECT * FROM accounts
WHERE id In (
    SELECT account_id FROM channel_subscribers
    GROUP BY account_id having count(channel_id)=2
    );
------------------------

-- متام پیام هایی که در گروه های مشرتک کاربر با شناسه hossein و ali وجود دارد را چاپ کنید.
SELECT m.content AS content
FROM messages m
         JOIN chats c ON m.chat_id = c.id
WHERE c.chat_type = 'group'
  AND m.chat_id IN (
    SELECT group_id
    FROM group_members
    WHERE group_id IN (
        SELECT group_id
        FROM group_members
        WHERE account_id IN (SELECT id FROM accounts WHERE username = 'Ali')
    )
      AND group_id IN (
        SELECT group_id
        FROM group_members
        WHERE account_id IN (SELECT id FROM accounts WHERE username = 'Hossein')
    )
);

--. میانگین تعداد کاراکرت پیام های ارسال شده به ربات با شناسه bot_first را چاپ کنید.
SELECT AVG(LENGTH(content)) AS average_message_length
FROM messages m join bots on m.chat_id=bots.id
         where  bots.id in (
             select id from bots
                       where username = 'bot_first'
             )

