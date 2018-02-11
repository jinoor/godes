using UnityEngine;
using System;
using System.Collections.Generic;
using System.Net;
using System.Net.Sockets;
using System.Threading;
using MsgPack.Serialization;
using MsgPack;

public class NetClient 
{
    private static NetClient instance_ = null;

    private string ip_ = "127.0.0.1";
    private int port_ = 4313;

    private TcpClient tcpClient_ = new TcpClient();
    private NetworkStream socketStream_;
    private MessagePackSerializer<List<object>> serializer_ = MessagePackSerializer.Get<List<object>>();
    private List<object> deserializedObject_ = new List<object>();

    public static NetClient Instance
    {
        get
        {
            if (instance_ == null)
                instance_ = new NetClient();
            return instance_;
        }
    }

    public void Connect()
    {
        try
        {
            IPAddress ipAddr = IPAddress.Parse(ip_);
            tcpClient_.Connect(new IPEndPoint(ipAddr, port_));
            Debug.Log("连接服务器成功");

            socketStream_ = tcpClient_.GetStream();
            var t = new Thread(Receive);
            t.Start();
        }
        catch
        {
            Debug.Log("连接服务器失败！");
        }
    }

    public void Send(byte[] buffer)
    {
        if (tcpClient_.Connected)
        {
            //Debug.Log("发送消息长度 >>> " + buffer.Length);
            socketStream_.Write(buffer, 0, buffer.Length);
            socketStream_.Flush();
        }
    }

    public void Receive()
    {
        while (tcpClient_.Connected && socketStream_.CanRead)
        {
            deserializedObject_.Clear();
            deserializedObject_ = serializer_.Unpack(socketStream_);

            List<MessagePackObject> recvObject = new List<MessagePackObject>();
            foreach (MessagePackObject o in deserializedObject_)
            {
                recvObject.Add(o);
            }

            NetTimer.dataRecv_.Add(recvObject);
        }
    }

    public void Close()
    {
        Debug.Log("关闭连接");
        tcpClient_.Close();
    }


}
