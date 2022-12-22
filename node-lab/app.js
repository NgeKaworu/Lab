/*
 * @Author: fuRan NgeKaworu@gmail.com
 * @Date: 2020-12-22 09:27:48
 * @LastEditors: fuRan NgeKaworu@gmail.com
 * @LastEditTime: 2022-12-16 16:27:48
 * @FilePath: /Lab/node-lab/app.js
 * @Description: 
 * 
 * Copyright (c) 2022 by fuRan NgeKaworu@gmail.com, All Rights Reserved. 
 */
const { json } = require('body-parser');
const { Server } = require('ws');

const sockserver = new Server({ port: 8888 });
sockserver.on('connection', (ws) => {
    console.log('New client connected!');
    ws.on('close', () => console.log('Client has disconnected!'));
    ws.on('message', (message) => {
        let body;
        try {
            body = JSON.parse(message)
        } catch { }

        if (body?.pointName) {
            sockserver.clients?.forEach(
                client =>
                    client.send(JSON.stringify({
                        type: "go_standby",
                        payload: {
                            mapId: "1",
                            pointName: body.pointName,
                            action: "robot_go_to_work"
                        }
                    }))
            )
        }
    })

});

