using UnityEngine;
using System.Collections;

public class Control : MonoBehaviour {

    private Entity entity_;
    private float lastMoveTime_;

	// Use this for initialization
	void Start () {
        entity_ = GetComponent<Entity>();
	
	}
	
	// Update is called once per frame
	void Update () {
	    InputKeyTest();
	}

    void InputKeyTest()
    {
        if (Input.GetKey(KeyCode.W))
        {
            entity_.MoveUp(1);
            SendMoveAsk();
        }
        else if (Input.GetKey(KeyCode.S))
        {
            entity_.MoveDown(1);
            SendMoveAsk();
        }
        else if (Input.GetKey(KeyCode.A))
        {
            entity_.MoveLeft(1);
            SendMoveAsk();
        }
        else if (Input.GetKey(KeyCode.D))
        {
            entity_.MoveRight(1);
            SendMoveAsk();
        }
        
        if (Input.GetKeyDown(KeyCode.Space))
        {
            NetSend.Instance.SendPutBombAsk((int)entity_.transform.position.x, (int)entity_.transform.position.z);
        }
    }

    void SendMoveAsk()
    {
        if (Time.realtimeSinceStartup - lastMoveTime_ > 0.05)
        {
            lastMoveTime_ = Time.realtimeSinceStartup;
            NetSend.Instance.SendUserMoveAsk((int)transform.position.x, (int)transform.position.z);
        }
    }
}
