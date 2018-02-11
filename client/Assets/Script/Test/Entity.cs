using UnityEngine;
using System.Collections;

static class ENTITY_TYPE
{
    public const int USER = 1;
    public const int MONSTER = 2;
}

public class Entity : MonoBehaviour {

    private Vector3 pos_ = new Vector3(0, 0, 0);
    private Vector3 scale_ = new Vector3(10, 10, 10);
    private int type_ = 0;
    private int id_ = 0;
    //private int fullHp_ = 10;
    private int currHp_ = 10;

	// Use this for initialization
	void Start () {
	
	}
	
	// Update is called once per frame
	void Update () {
	
	}

    public void Init(int type, int id, int x, int y)
    {
        transform.localScale = scale_;
        type_ = type;
        id_ = id;
        pos_.x = x;
        pos_.z = y;
        transform.position = pos_;
    }

    public void UpdatePos(int x, int y)
    {
        pos_.x = x;
        pos_.z = y;
        transform.position = pos_;
    }

    public void MoveUp(int dev)
    {
        pos_.z += dev;
        transform.position = pos_;
    }

    public void MoveDown(int dev)
    {
        
        pos_.z -= dev;
        transform.position = pos_;
    }

    public void MoveLeft(int dev)
    {
        
        pos_.x -= dev;
        transform.position = pos_;
    }

    public void MoveRight(int dev)
    {
        
        pos_.x += dev;
        transform.position = pos_;
    }

    public void Destroy()
    {
        Destroy(gameObject);
    }

    void OnMouseDown()
    {
        if (type_ == ENTITY_TYPE.MONSTER)
        {
            NetSend.Instance.SendAttackAsk(id_);
        }
    }

    public void UpdateHp(int hp)
    {
        scale_.y = hp;
        transform.localScale = scale_;

        currHp_ = hp;
    }
}
