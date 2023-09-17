USE `yubisuma_api`;

SET NAMES utf8mb4;

INSERT INTO
`animals` (
`id`,
`name`,
`numFinger`,
`skill`,
`skillDesc`,
`handUrl`,
`upFingerUrl`,
`downFingerUrl`
    )
VALUES (
        UUID(),
        "人間",
        5,
        "さいみんじゅつ",
        "あいてを せんのうして つぎの あいての たてるゆびを いっぽんに しちゃうぞ！",
        "handUrl_for_人間",
        "upFingerUrl_for_人間",
        "downFingerUrl_for_人間"
    ), (
        UUID(),
        "恐竜",
        2,
        "ほうこう",
        "びっくりさせて あいての じょがいしていた ゆびを いっぽん ふっかつさせるぞ！",
        "handUrl_for_恐竜",
        "upFingerUrl_for_恐竜",
        "downFingerUrl_for_恐竜"
    ), (
        UUID(),
        "たこ",
        8,
        "すみをはく",
        "すみで おたがいの ゆびを かくしちゃうぞ！たいへん！",
        "handUrl_for_たこ",
        "upFingerUrl_for_たこ",
        "downFingerUrl_for_たこ"
    );