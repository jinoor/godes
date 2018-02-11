using UnityEngine;
using System;
using System.Collections;
using System.Collections.Generic;
using MsgPack;

public class NetTimer : MonoBehaviour {

    public static List<List<MessagePackObject>> dataRecv_ = new List<List<MessagePackObject>>();

	// Use this for initialization
	void Start () {
	
	}
	
	// Update is called once per frame
	void Update () 
    {
        if (dataRecv_.Count > 0)
        {
            try
            {
                NetRecv.Instance.Recv(dataRecv_[0]);
            }
            catch (Exception ex)
            {
                Debug.Log("Exception from deal with msg: " + ex.Message + " stack: " + ex.StackTrace);
                if (dataRecv_[0].Count > 0)
                {
                    Debug.Log("The error proto: " + dataRecv_[0][0]);
                }
            }
            finally
            {
                dataRecv_.RemoveAt(0);
            }
        }
	
	}
}
